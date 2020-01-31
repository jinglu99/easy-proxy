package actions

import (
	"fmt"
	"github.com/jingleWang/easy-proxy/src/config"
	"github.com/jingleWang/easy-proxy/src/network"
	"github.com/urfave/cli/v2"
	"strconv"
)

func ListNetwork(c *cli.Context) error {
	network.PrintAllNetworkServices()
	return nil
}

func GetNetwork(c *cli.Context) error {
	ns := config.GetNetwork()
	if ns == "" {
		fmt.Println("default network is not set, you can use list networks and config network set <network> to set.")
	} else {
		fmt.Println(ns)
		fmt.Println("you can use config network set <network> to change it.")
	}
	return nil
}

func SetNetwork(c *cli.Context) error {
	if c.NArg() < 1 {
		fmt.Println("required an argument.")
		return nil
	}

	arg := c.Args().Get(0)
	if i, err := strconv.Atoi(arg); err != nil {
		setNetwrokByName(arg)
	} else {
		setNetwrokByIndex(i)
	}
	return nil
}

func setNetwrokByName(name string) {
	for _, n := range network.ListNetworkServices() {
		if n == name {
			config.SetNetwork(name)
			return
		}
	}

	fmt.Printf("there is not a network service called %s\n", name)
}

func setNetwrokByIndex(index int) {
	networks := network.ListNetworkServices()
	if index >= 0 && index < len(networks) {
		config.SetNetwork(networks[index])
		return
	}

	fmt.Printf("the index %d is out of bound\n", index)
}

