package bingo

import (
	"github.com/spf13/viper"
)

// Config 配置
type Config struct {
	*viper.Viper
	Server *ServerConfig
	Mysql  *MysqlConfig
	Mongo  *MongoConfig
	Custom UserConfig
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
	DB       string
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
	v.SetConfigName("config")
	v.AddConfigPath("conf")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := &Config{
		Viper:  v,
		Server: &ServerConfig{},
		Mysql:  &MysqlConfig{},
		Mongo:  &MongoConfig{},
	}

	conf.Section("Server", conf.Server)
	conf.Section("Mysql", conf.Mysql)
	conf.Section("Mongo", conf.Mongo)

	return conf
}

var (
	sectionMap = make(map[string]interface{})
)

// Section 读取指定配置
func (c *Config) Section(key string, value interface{}) error {
	err := c.Viper.UnmarshalKey(key, value)

	if err != nil {
		return err
	}

	if _, ok := sectionMap[key]; !ok {
		sectionMap[key] = value
	}

	return nil
}

// Reload 重新加载
func (c *Config) Reload() error {
	for k, v := range sectionMap {
		if err := c.Section(k, v); err != nil {
			return err
		}
	}

	return nil
}
