package config

const networkConf string = "network"
func GetNetwork() string {
	return conf.GetString(networkConf)
}

func SetNetwork(network string) {
	conf.Set(networkConf, network)
	conf.WriteConfig()
}
