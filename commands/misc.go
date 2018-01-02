package commands

import (
	"fmt"

	"github.com/arteev/dsql/rdb"
	"github.com/arteev/dsql/repository"
	"github.com/urfave/cli"
)

//GetCommandsMisc define cli command: env
func GetCommandsMisc() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:  "env",
			Usage: "print dsql environment(variables,options) information",
			Action: func(ctx *cli.Context) {
				fmt.Println("REPOSITORY:", repository.GetRepositoryFile(), " DEFAULT:", repository.IsDefault())
				fmt.Println("ENGINES:", rdb.KnownEngine)
			},
		},
	}
}
