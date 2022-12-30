package store

type Config struct {
	DatabaseURL string `toml: "database_url"`
}

func Newconfig() *Config {
	return &Config{}
}
