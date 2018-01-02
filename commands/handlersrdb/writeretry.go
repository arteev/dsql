package handlersrdb

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arteev/dsql/db"
	"github.com/arteev/dsql/rdb/action"
	"github.com/arteev/logger"
)

func WriteRetryFile(dbs []db.Database, ctx *action.Context) error {
	var err error
	retryName := ctx.GetDef("retryname", "").(string)
	defaultRetry := "alias.retry"
	if retryName == "" {
		dir, err := os.Getwd()
		if err != nil {
			retryName = filepath.Join(".", defaultRetry)
		} else {
			retryName = filepath.Join(dir, defaultRetry)
		}
		ctx.Set("retryname", retryName)
	}

	if retryName != "" {
		if err := os.Remove(retryName); err != nil {
			logger.Error.Println(err)
		}
	}

	var fretry *os.File
	defer func() {
		if fretry != nil {
			fretry.Close()
		}
	}()
	for _, d := range dbs {
		dbCtx := ctx.Get("context" + d.Code).(*action.Context)
		if !dbCtx.Get("success").(bool) {
			if fretry == nil {
				fretry, err = os.Create(retryName)
				if err != nil {
					return fmt.Errorf("Could not write retry file: %v", err)
				}
			}
			fretry.WriteString(d.Code + "\n")
		}
	}
	if !ctx.GetDef("silent", false).(bool) && fretry != nil {
		fmt.Printf("There were mistakes. Use to repeat -d @%s\n", retryName)
	}
	return nil
}
