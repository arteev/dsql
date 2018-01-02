package app

import (
	"fmt"

	"github.com/arteev/dsql/commands"
	"github.com/arteev/logger"
	"github.com/urfave/cli"

	"os"

	"github.com/arteev/dsql/repository"
)

//A Application cli
type Application struct {
	verbose int //debug level
	cli     *cli.App
}

//New cli application
func newApp() *Application {
	a := &Application{
		cli: newcli(),
	}
	return a.globalFlags().
		beforeAction().
		defineCommands().
		defaultAction()
}

//newcli. Create new instance of cli.App with commands, flags, etc.
func newcli() *cli.App {
	c := cli.NewApp()
	c.Version = fmt.Sprintf("%s (%s) githash:%s", Version, DateBuild, GitHash)
	c.Description = "It's the tool for the simultaneous execution  SQL statements in multiple databases"
	c.Copyright = "MIT Licence"
	c.Name = AppName
	c.Author = Authors
	return c
}

func (a *Application) globalFlags() *Application {
	a.cli.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "verbose",
			Usage: "set sprecific debug output level",
		},
		cli.StringFlag{
			Name:  "repo,r",
			Usage: "select repository. default repository.sqlite",
		},
		cli.BoolFlag{
			Name:  "stat,s",
			Usage: "show common statistic for all query and etc.",
		},
		cli.BoolFlag{
			Name:  "statquery,sq",
			Usage: "show statistic for each query",
		},
		cli.BoolFlag{
			Name:  "silent",
			Usage: "Hide runtime error",
		},
	}
	return a
}

func (a *Application) defineCommands() *Application {
	arr := [][]cli.Command{
		commands.GetCommandsListDB(),
		commands.GetCommandsParams(),
		commands.GetCommandsDBS(),
		commands.GetCommandsMisc(),
	}
	a.cli.Commands = []cli.Command{}
	for _, item := range arr {
		a.cli.Commands = append(a.cli.Commands, item...)
	}
	return a
}

func (a *Application) beforeAction() *Application {
	a.cli.Before = func(ctx *cli.Context) error {
		a.verbose = ctx.GlobalInt("verbose")
		logger.InitToConsole(logger.Level(a.verbose))
		logger.Info.Println("Verbose level:", logger.CurrentLevel)

		var rfile string
		if ctx.GlobalIsSet("repo") {
			rfile = ctx.GlobalString("repo")
		} else if ctx.GlobalIsSet("r") {
			rfile = ctx.GlobalString("r")
		}
		repository.SetRepositoryFile(rfile)

		logger.Info.Println("Verbose level:", logger.CurrentLevel)
		logger.Info.Println("Repository location:", repository.GetRepositoryFile(), "Default:", repository.IsDefault())
		return nil
	}
	return a
}

//Run - entry point to the cli app.
func Run() error {
	logger.InitToConsole(logger.LevelTrace)

	repository.Init()
	defer repository.Done()

	a := newApp()
	return a.cli.Run(os.Args)
}

func (a *Application) defaultAction() *Application {
	return a
}
