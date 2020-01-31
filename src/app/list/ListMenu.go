package list

import (
	"github.com/jingleWang/easy-proxy/src/actions"
	"github.com/urfave/cli/v2"
)

func Commands() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "get the list of rules/networks",
		Subcommands: []*cli.Command{
			{
				Name:   "networks",
				Usage:  "list all network service for this machine.",
				Action: actions.ListNetwork,
			},
			{
				Name:   "rules",
				Usage:  "list all rules that added.",
				Action: actions.ListRules,
			},
		},
	}
}
