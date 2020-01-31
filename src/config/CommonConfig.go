package config

const (
	portConf = "port"
)

func SetPort(port int) {
	conf.Set(portConf, port)
	conf.WriteConfig()
}

func GetPort() int {
	if port := conf.GetInt(portConf); port <= 0 || port > 65535 {
		return 2532 //default port
	} else {
		return port
	}

}
