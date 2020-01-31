package app

import (
	"github.com/jingleWang/easy-proxy/src/app/add"
	"github.com/jingleWang/easy-proxy/src/app/config"
	"github.com/jingleWang/easy-proxy/src/app/list"
	"github.com/jingleWang/easy-proxy/src/app/remove"
	"github.com/jingleWang/easy-proxy/src/app/start"
	"github.com/urfave/cli/v2"
)

func MainMenu() cli.App {
	return cli.App{
		Name:                 "easyProxy",
		Usage:                "An proxy for backend developer to transfer request to local environment.",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			start.Commands(),
			list.Commands(),
			config.Commands(),
			add.Commands(),
			remove.Commands(),
		},
	}
}
