package repository

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"runtime"

	"github.com/arteev/logger"

	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

//RepositoryFile - current file of repository
var (
	repositoryFile = "repository.sqlite"
	isDefault      = true
	mustRemove     = false
	tmpFile        string
)

func Init() {
	search()
}

func Done() {
	if mustRemove {
		os.Remove(tmpFile)
	}
}

func search() {
	//workdir
	if _, err := os.Stat(repositoryFile); err == nil {
		return
	}
	var cfgLocation string
	//appdata | ~/.config
	if u, err := user.Current(); err == nil {
		if runtime.GOOS == "windows" {
			cfgLocation = filepath.Join(os.Getenv("APPDATA"), "dsql", repositoryFile)
		} else {
			cfgLocation = filepath.Join(u.HomeDir, ".config", "dsql", repositoryFile)
		}
		if _, err := os.Stat(cfgLocation); err == nil {
			repositoryFile = cfgLocation
			return
		}
	}
	//folder dsql
	absPath, _ := filepath.Abs(path.Dir(os.Args[0]))
	inAppLocation := path.Join(absPath, repositoryFile)
	if _, err := os.Stat(inAppLocation); err == nil {
		repositoryFile = inAppLocation
		return
	}
	if cfgLocation != "" {
		repositoryFile = cfgLocation
	}
}

//SetRepositoryFile - set new location repository file
func SetRepositoryFile(filename string) error {
	if filename == "" {
		return nil
	}
	if !isDefault {
		return fmt.Errorf("can't twice change repository file ")
	}
	if filename != "" {
		isDefault = false
		repositoryFile = filename
	}
	return nil
}

//GetRepositoryFile - get current location repository file
func GetRepositoryFile() string {
	return repositoryFile
}

//IsDefault returns location repository file is default
func IsDefault() bool {
	return isDefault
}

//PrepareLocation - make directories for repository files
func PrepareLocation() error {
	if url, err := url.Parse(repositoryFile); err == nil && url.Scheme != "" {
		tmp, err := ioutil.TempFile("", "rep.sqlite3")
		if err != nil {
			return err
		}
		tmpFile = tmp.Name()
		mustRemove = true
		defer tmp.Close()
		resp, err := http.Get(repositoryFile)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if _, err := io.Copy(tmp, resp.Body); err != nil {
			return err
		}
		repositoryFile = "file:///" + tmp.Name() + "?" + url.RawQuery
		logger.Info.Println("Repository temp:", repositoryFile, "on disk:", tmpFile)
		return nil
	}

	dir := filepath.Dir(repositoryFile)
	if dir == "" || dir == "." {
		return nil
	}
	perm := 0700
	if err := os.MkdirAll(dir, os.FileMode(perm)); err != nil {
		return err
	}
	return nil
}
