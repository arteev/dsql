package main

import (
	"github.com/arteev/dsql/app"
	"github.com/arteev/fmttab"
	"github.com/arteev/logger"
)

func main() {
	fmttab.Trimend = ">"
	if err := app.Run(); err != nil {
		logger.Error.Fatal(err)
	}
}
