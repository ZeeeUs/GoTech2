package store

type Config struct {
	DatabaseURL string
}

func NewDbConfig() *Config {
	return &Config{
		DatabaseURL: "/home/zeus/test/one.json",
	}
}
