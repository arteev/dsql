package main

import (
	//"fmt"

	"github.com/arteev/dsql/app"
	//"github.com/arteev/logger"
	"github.com/arteev/fmttab"
	"github.com/arteev/dsql/repofile"
)

func main() {
	/*defer func() {
		e := recover()
		if e != nil {
			if logger.CurrentLevel < logger.LevelError {
				fmt.Println(e)
			}
			logger.Error.Println(e)
		}
	}()*/
	fmttab.Trimend = ">"
	a := app.New()
	defer repofile.Done()
	if err := a.Run(); err != nil {
		panic(err)
	}


}
