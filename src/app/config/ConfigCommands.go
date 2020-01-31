package config

import (
	"github.com/jingleWang/easy-proxy/src/app/config/network"
	"github.com/jingleWang/easy-proxy/src/app/config/port"
	"github.com/urfave/cli/v2"
)

func Commands() *cli.Command {
	return &cli.Command{
		Name:    "config",
		Aliases: []string{"conf"},
		Usage:   "get or set default config",
		Subcommands: []*cli.Command{
			network.Commands(),
			port.Commands(),
		},
	}
}
