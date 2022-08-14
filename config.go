package bingo

import (
	"github.com/spf13/viper"
)

// Config 配置
type Config struct {
	Server *ServerConfig
	Mysql  *MysqlConfig
	Mongo  *MongoConfig
	Custom UserConfig
}

func NewConfig() *Config {
	return &Config{
		Server: &ServerConfig{
			Port:    8080,
			AppMode: "debug",
		},
		Mysql: &MysqlConfig{
			Host:     "127.0.0.1",
			Port:     3306,
			User:     "",
			Password: "",
			Database: "",
		},
	}
}

// UserConfig 用户自定义配置
type UserConfig map[interface{}]interface{}

// GetConfig 递归读取用户配置文件
func GetConfig(config UserConfig, prefix []string, index int) interface{} {
	key := prefix[index]
	if v, ok := config[key]; ok {
		if index == len(prefix)-1 {
			return v
		}

		index = index + 1
		if val, ok := v.(UserConfig); ok {
			return GetConfig(val, prefix, index)
		} else {
			return nil
		}
	}

	return nil
}

// ServerConfig 服务配置
type ServerConfig struct {
	Port    int
	AppMode string
	Html    string
}

// MysqlConfig mysql数据库配置
type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string
}

// MongoConfig mongo数据库配置
type MongoConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

// InitConfig 初始化配置文件
func InitConfig() *Config {
	v := viper.New()
	v.AddConfigPath("./")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := NewConfig()
	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}

	return conf
}
