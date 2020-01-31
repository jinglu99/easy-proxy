package network

import (
	"os/exec"
	"strconv"
)

const (
	autoDiscovery    = "autoDiscovery"
	pac              = "pac"
	webProxy         = "webProxy"
	securityWebProxy = "securityWebProxy"
	ftpProxy         = "ftpProxy"
	socksProxy       = "socksProxy"
	rtsp             = "rtsp"
	gopherProxy      = "gopherProxy"
)

var managers = map[string]*Manager{
	autoDiscovery: {
		stateCmd: "-setproxyautodiscovery",
		getCmd:   "-getproxyautodiscovery",
		setCmd:   "-setproxyautodiscovery",
	},
	pac: {
		stateCmd: "-setautoproxystate",
		getCmd:   "-getautoproxyurl",
		setCmd:   "-setautoproxyurl",
	},
	webProxy: {
		stateCmd: "-setwebproxystate",
		getCmd:   "-getwebproxy",
		setCmd:   "-setwebproxy",
	},
	securityWebProxy: {
		stateCmd: "-setsecurewebproxystate",
		getCmd:   "-getsecurewebproxy",
		setCmd:   "-setsecurewebproxy",
	},
	ftpProxy: {
		stateCmd: "-setftpproxystate",
		getCmd:   "-getftpproxy",
		setCmd:   "-setftpproxy",
	},
	socksProxy: {
		stateCmd: "-setsocksfirewallproxystate",
		getCmd:   "-getsocksfirewallproxy",
		setCmd:   "-setsocksfirewallproxy",
	},
	rtsp: {
		stateCmd: "-setstreamingproxystate",
		getCmd:   "-getstreamingproxy",
		setCmd:   "-setstreamingproxy",
	},
	gopherProxy: {
		stateCmd: "-setgopherproxystate",
		getCmd:   "-getgopherproxy",
		setCmd:   "-setgopherproxy",
	},
}

func switchOffAllProxy(network string) {
	for _, m := range managers {
		go m.Off(network)
	}
}

func setSocksProxy(network string, host string, port int) {
	managers[socksProxy].Set(network, host, strconv.Itoa(port))
}

type Manager struct {
	stateCmd string
	getCmd   string
	setCmd   string
}

func (m *Manager) On(network string) {
	exec.Command("networksetup", m.stateCmd, network, "on").Run()
}
func (m *Manager) Off(network string) {
	exec.Command("networksetup", m.stateCmd, network, "off").Run()
}
func (m *Manager) Set(network string, args ...string) {
	cmdArgs := []string {m.setCmd, network}
	cmdArgs = append(cmdArgs, args...)
	exec.Command("networksetup", cmdArgs...).Run()
}
func (m *Manager) Get(network string) string {
	return ""
}
