package main

import (
	"fmt"

	"github.com/arteev/dsql/app"
	"github.com/arteev/logger"
)

func main() {
	defer func() {
		e := recover()
		if e != nil {
			if logger.CurrentLevel < logger.LevelError {
				fmt.Println(e)
			}
			logger.Error.Println(e)
		}
	}()
	a := app.New()
	if err := a.Run(); err != nil {
		panic(err)
	}
}
