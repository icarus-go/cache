package config

import "pmo-test4.yz-intelligence.com/kit/cache/adapter"

// Config 配置文件
type Config struct {
	Scheme   string
	Host     string
	Port     string
	Password string
	Adapter  adapter.Adapter
}

// New 配置文件实例
func New(adapter adapter.Adapter, scheme, host, port, password string) *Config {
	return &Config{
		scheme,
		host,
		port,
		password,
		adapter,
	}
}
