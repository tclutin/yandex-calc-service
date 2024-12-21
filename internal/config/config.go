package config

type Config struct {
	Host string
	Port string
}

func New() *Config {
	return &Config{}
}
