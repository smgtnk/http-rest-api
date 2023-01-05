package store

type Config struct {
	DatabaseURL string `toml: "database_url"`
}

func Newconfig() *Config {
	return &Config{
		// DatabaseURL: "sqlserver://sa:123@localhost:51000?database=restapi_dev&connection+timeout=30",
		DatabaseURL: "host=localhost user=postgres password=123 dbname=restapi_dev sslmode=disable",
	}
}
