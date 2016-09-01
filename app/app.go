package app

import (
	"github.com/arteev/dsql/commands"
	"github.com/arteev/logger"
	"github.com/urfave/cli"

	"os"

	"github.com/arteev/dsql/repofile"
)

const (
	// Version of application (for installation outside)
	Version = "1.3"
)

//A Application cli
type Application struct {
	verbose int //debug level
	cli     *cli.App
}

//New cli application
func New() *Application {
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
	c.Version = Version
	c.Copyright = "MIT Licence"
	c.Name = "dsql"
	c.Usage = "It is a tool for the simultaneous execution of multiple SQL statements in the database"
	c.Author = "Arteev Aleksey"
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
	a.cli.Commands = []cli.Command{}
	a.cli.Commands = append(a.cli.Commands,
		commands.GetCommandsListDB()...,
	)
	a.cli.Commands = append(a.cli.Commands,
		commands.GetCommandsParams()...,
	)
	a.cli.Commands = append(a.cli.Commands,
		commands.GetCommandsDBS()...,
	)
	a.cli.Commands = append(a.cli.Commands,
		commands.GetCommandsConfig()...,
	)
	return a
}

func (a *Application) beforeAction() *Application {
	a.cli.Before = func(ctx *cli.Context) error {
		a.verbose = ctx.GlobalInt("verbose")
		logger.InitToConsole(logger.Level(a.verbose))
		logger.Info.Println("Verbose level:", logger.CurrentLevel)
		if ctx.GlobalIsSet("repo") {
			repofile.SetRepositoryFile(ctx.GlobalString("repo"))
		} else if ctx.GlobalIsSet("r") {
			repofile.SetRepositoryFile(ctx.GlobalString("r"))
		}

		logger.Info.Println("Verbose level:", logger.CurrentLevel)
		logger.Info.Println("Repository location:", repofile.GetRepositoryFile(), "Default:", repofile.IsDefault())
		return nil
	}
	return a
}

//Run - entry point to the cli app.
func (a *Application) Run() error {
	return a.cli.Run(os.Args)
}

func (a *Application) defaultAction() *Application {
	return a
}
