package apiserver

type Config struct {
	port  string
	alaDb string
}

func NewConfig() *Config {
	return &Config{
		port: ":8080",
		alaDb: "db.txt",
	}
}
