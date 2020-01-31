package port

import (
	"github.com/jingleWang/easy-proxy/src/actions"
	"github.com/urfave/cli/v2"
)

func Commands() *cli.Command {
	return &cli.Command{
		Name:    "port",
		Description:"get or set default port",
		Usage:   "port [get|set] [argument]",
		Subcommands: []*cli.Command{
			{
				Name:  "set",
				Description:"set default port.",
				Usage: "set <port>",
				Action: actions.SetPort,
			},
			{
				Name:  "get",
				Description:"get default port",
				Action: actions.GetPort,
			},
		},
		Action:actions.GetPort,
	}
}