package actions

import (
	"fmt"
	"github.com/armon/go-socks5"
	"github.com/jingleWang/easy-proxy/src/config"
	"github.com/jingleWang/easy-proxy/src/network"
	"github.com/jingleWang/easy-proxy/src/rewriter"
	"github.com/jingleWang/easy-proxy/src/rule"
	"github.com/urfave/cli/v2"
	"github.com/ztrue/shutdown"
	"os"
	"strconv"
	"syscall"
)

func StartProxy(c *cli.Context) error {
	setDefaultNetworkServiceIfNotSet()

	port := config.GetPort()

	if pStr := c.String("port"); pStr != "" {
		if p, err := strconv.Atoi(pStr); err == nil {
			port = p
		}
	}

	go start(port, c.String("rule"))

	fmt.Printf("proxy started on port: %d\n", port)
	fmt.Printf("Press ctrl + c to exit...")

	network.SetSock5Proxy(config.GetNetwork(), "localhost", port)

	addShutdownHook()
	return nil
}

func setDefaultNetworkServiceIfNotSet() {
	if config.GetNetwork() != "" {
		return
	}

	network.PrintAllNetworkServices()
	var ns int
	networks := network.ListNetworkServices()
	fmt.Print("please type the index of network service you are using: ")
	fmt.Scanf("%d", &ns)
	if ns >= 0 && ns < len(networks) {
		config.SetNetwork(networks[ns])
		fmt.Printf("config default network service: %s\n", config.GetNetwork())
	}
}

func start(port int, r string) {

	var conf *socks5.Config
	if r == "" {
		conf = &socks5.Config{Rewriter: &rewriter.EasyRewriter{Rules: config.GetRules()}}
	} else {
		conf = &socks5.Config{Rewriter: &rewriter.EasyRewriter{Rules: &[]rule.Rule{rule.MustNewRule(r)}}}
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost
	if err := server.ListenAndServe("tcp", "127.0.0.1:"+strconv.Itoa(port)); err != nil {
		panic(err)
	}
}

func addShutdownHook() {
	shutdown.AddWithParam(func(sig os.Signal) {
		network.SwitchOffSocksProxy(config.GetNetwork())
		print("\nProxy Stopped...\n")
		os.Exit(9)
	})
	shutdown.Listen(syscall.SIGINT)
}
