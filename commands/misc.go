package commands

import (
	"fmt"

	"github.com/arteev/dsql/rdb"
	"github.com/arteev/dsql/repository"
	"github.com/urfave/cli"
)

func init() {
	Register([]cli.Command{
		cli.Command{
			Name:  "env",
			Usage: "print dsql environment(variables,options) information",
			Action: func(ctx *cli.Context) {
				fmt.Println("Environment variables.")
				fmt.Printf("%s : %q\n", repository.ENDSQLREPO, repository.EnvDSQLRepo)
				fmt.Println("DSQL variables.")
				fmt.Printf("REPOSITORY: %s DEFAULT: %v\n", repository.GetRepositoryFile(), repository.IsDefault())
				fmt.Printf("ENGINES: %v\n", rdb.KnownEngine)
			},
		},
	})
}
