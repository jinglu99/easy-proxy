package network

import (
	"fmt"
	"os/exec"
	"strings"
)

func ListNetworkServices() []string {
	r := exec.Command("networksetup" , "-listallnetworkservices")
	out, _ := r.CombinedOutput()

	services := strings.Split(string(out), "\n")
	return services[1:len(services) - 1]
}

func PrintAllNetworkServices() {
	for i, v := range ListNetworkServices() {
		fmt.Printf("(%d).%s\n", i, v)
	}
}

func SetSock5Proxy(network string, host string, port int) error {
	switchOffAllProxy(network)
	setSocksProxy(network, host, port)
	SwitchOnSocksProxy(network)
	return nil
}

func SwitchOffSocksProxy(network string)  {
	managers[socksProxy].Off(network)
}

func SwitchOnSocksProxy(network string)  {
	managers[socksProxy].On(network)
}
