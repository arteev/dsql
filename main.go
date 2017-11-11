package main

import (
	"github.com/arteev/dsql/app"
	"github.com/arteev/dsql/repofile"
	"github.com/arteev/fmttab"
	"github.com/arteev/logger"
)

func main() {
	fmttab.Trimend = ">"
	defer repofile.Done()
	if err := app.Run(); err != nil {
		logger.Error.Fatal(err)
	}
}
