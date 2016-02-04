package sqlcommand

import (
	"io/ioutil"
	"strings"

	"github.com/arteev/logger"
)

//A SQLCommand this command sql for remote database
type SQLCommand struct {
	Script string
	Params []string
}

//New *SQLCommand
func New(text string, params []string) *SQLCommand {
	//load from file
	if strings.Index(text, "@") == 0 {
		file := string(text[1:])
		logger.Debug.Println("SQL from file:", file)
		sqltext, e := ioutil.ReadFile(file)
		if e != nil {
			panic(e)
		}

		return &SQLCommand{
			Script: string(sqltext),
			Params: params,
		}
	}

	return &SQLCommand{
		Script: text,
		Params: params,
	}
}
