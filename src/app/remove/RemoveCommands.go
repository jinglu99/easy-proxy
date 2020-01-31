package remove

import (
	"github.com/jingleWang/easy-proxy/src/actions"
	"github.com/urfave/cli/v2"
)

//remove menu
func Commands() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"rm"},
		Description: "remove rule",
		Usage:   "remove <ruleIndex>, remove rule by index refer to <list rules>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "all",
				Usage:   "remove all rules",
				Aliases: []string{"a"},
			},
		},
		Action: actions.RemoveRule,
	}
}
