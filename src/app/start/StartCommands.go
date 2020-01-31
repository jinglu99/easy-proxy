package start

import (
	"github.com/jingleWang/easy-proxy/src/actions"
	"github.com/urfave/cli/v2"
)

func Commands() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "start proxy",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "rule",
				Aliases: []string{"r"},
				Usage:   "start proxy for given rule, if not provided, proxy all rules added.",
			}, &cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "start proxy on given port, default in 2532",
			},
		},
		Action: actions.StartProxy,
	}
}
