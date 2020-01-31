package actions

import (
	"fmt"
	"github.com/jingleWang/easy-proxy/src/config"
	"github.com/urfave/cli/v2"
	"strconv"
)

func SetPort(c *cli.Context) error {
	if c.NArg() < 1 {
		fmt.Println("must provide port")
	}

	portStr := c.Args().Get(0)
	if port, err := strconv.Atoi(portStr); err != nil {
		fmt.Println("required an integer argument.")
	} else if port <= 0 || port > 65535 {
		fmt.Println("expected port must between 1 and 65535.")
	} else {
		config.SetPort(port)
	}

	return nil
}

func GetPort(c *cli.Context) error {
	fmt.Println(strconv.Itoa(config.GetPort()))
	return nil
}
