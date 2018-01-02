package commands

import (
	"github.com/urfave/cli"
)

var cmds = make([]cli.Command, 0)

//Register commands cmd
func Register(cmd []cli.Command) {
	cmds = append(cmds, cmd[:]...)
}

//Get returns registred commands
func Get() []cli.Command {
	return cmds
}
