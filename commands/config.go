package commands

import (
	"fmt"

	"github.com/arteev/dsql/rdb"
	"github.com/arteev/dsql/repofile"
	"github.com/arteev/logger"
	"github.com/urfave/cli"
)

//GetCommandsConfig define cli command config
func GetCommandsConfig() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:  "config",
			Usage: "managment engines",
			Subcommands: []cli.Command{
				cli.Command{
					Name:  "engines",
					Usage: "list known engines",
					Action: func(ctx *cli.Context) {
                        logger.Trace.Println("command engines list")
						defer logger.Trace.Println("command engines list done")
						for _, e := range rdb.KnownEngine {
							fmt.Println(e)
						}
					},
				},
                cli.Command{
					Name:  "location",
					Usage: "type location repository",
					Action: func(ctx *cli.Context) {
                        logger.Trace.Println("command location")
                        fmt.Println("repository:",repofile.GetRepositoryFile()," default:",repofile.IsDefault())
						defer logger.Trace.Println("command location done")
						
					},
				},
			}, //subcommands
		},
	}
}
