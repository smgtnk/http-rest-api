package apiserver

import "github.com/smgtnk/http-rest-api/store"

type Config struct {
	BindAddr string `toml: "bind_addr"`
	LogLevel string `toml: "log_level"`
	Store    *store.Config
}

func Newconfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.Newconfig(),
	}
}
