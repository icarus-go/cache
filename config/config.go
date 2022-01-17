package config

import "github.com/icarus-go/cache/adapter"

// Config 配置文件
type Config struct {
	Scheme   string          `json:"scheme" yaml:"scheme"`
	Host     string          `json:"host" yaml:"host"`
	Port     string          `json:"port" yaml:"port"`
	Password string          `json:"password" yaml:"password"`
	Adapter  adapter.Adapter `json:"adapter" yaml:"adapter"`
	DB       int             `json:"db" yaml:"db"`
}

// New 配置文件实例
func New(adapter adapter.Adapter, scheme, host, port, password string, db int) *Config {
	return &Config{
		scheme,
		host,
		port,
		password,
		adapter,
		db,
	}
}

//DefaultNew 默认使用DB0
func DefaultNew(adapter adapter.Adapter, scheme, host, port, password string) *Config {
	return &Config{
		scheme,
		host,
		port,
		password,
		adapter,
		0,
	}
}
