package config

import "github.com/ZeeeUs/GoTech2/develop/11_HTTPServer/store"

type Config struct {
	Port string
	Store *store.Config
}

func NewDefConfig() *Config {
	return &Config{
		Port: ":8080",
		Store: store.NewDbConfig(),
	}
}
