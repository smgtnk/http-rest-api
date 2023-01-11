package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

func Newconfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		//DatabaseURL: "host=localhost user=postgres password=123 dbname=restapi_test sslmode=disable",
	}
}
