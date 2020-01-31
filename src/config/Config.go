package config

import (
	"github.com/spf13/viper"
	"os"
)

var (
	userHome = os.Getenv("HOME")
	rootPath = userHome + "/.easyproxy/"
)
var conf *viper.Viper = viper.New()

func init() {
	mkDirIfNotExist(rootPath)
	conf.SetConfigName("config")
	conf.AddConfigPath(rootPath)
	conf.SetConfigType("yaml")
	conf.SafeWriteConfig()
	conf.ReadInConfig()
}

func mkDirIfNotExist(path string) error {
	if isExist, err := pathExists(path); err != nil {
		return err
	} else {
		if !isExist {
			os.MkdirAll(path, os.ModePerm)
		}
		return nil
	}
}

// 判断文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
