package add

import (
	"github.com/jingleWang/easy-proxy/src/actions"
	"github.com/urfave/cli/v2"
)

func Commands() *cli.Command {
	return &cli.Command{
		Name:  "add",
		Description:"add rule",
		Usage: "add <pattern> <dest_host:port>, add new rule",
		Action: actions.AddRules,
	}
}
