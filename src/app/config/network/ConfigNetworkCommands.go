package network

import (
	"github.com/jingleWang/easy-proxy/src/actions"
	"github.com/urfave/cli/v2"
)

func Commands() *cli.Command {
	return &cli.Command{
		Name:    "network",
		Description:"get or set default network service config",
		Usage:   "network [get|set] [argument]",
		Subcommands: []*cli.Command{
			{
				Name:  "set",
				Description:"set network by name or index.",
				Usage: "set <network>",
				Action: actions.SetNetwork,
			},
			{
				Name:  "get",
				Description: "get default network service config",
				Action: actions.GetNetwork,
			},
		},
		Action:actions.GetNetwork,
	}
}
