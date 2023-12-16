package config

import "os"

var Config AppConfig

type AppConfig struct {
	port string
}

func LoadConfig() AppConfig {
	Config = AppConfig{
		port: os.Getenv("GOLFS_PORT"),
	}
	return Config
}

func (c AppConfig) Port() string {
	if c.port == "" {
		return "8989"
	} else {
		return c.port
	}
}
