package bingo

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

// InitConfig 初始化配置文件
func InitConfig(dir, filename string) error {
	if dir == "" || filename == "" {
		return fmt.Errorf("incorrect configuration settings")
	}

	viper.AddConfigPath(dir)
	if strings.Contains(filename, ".") {
		arr := strings.Split(filename, ".")
		viper.SetConfigName(arr[0])
		viper.SetConfigType(arr[1])
	} else {
		viper.SetConfigName(filename)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("%s file not found \n", filename)
		} else {
			return fmt.Errorf("Fatal error config file: %s \n", err)
		}
	}

	viper.WatchConfig()
	return nil
}
