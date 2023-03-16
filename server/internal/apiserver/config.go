package apiserver

import "github.com/FFH255/CW-2023/internal/hepler"

type Config struct {
	Port     string
	LogLever string
}

func NewConfig() *Config {
	return &Config{
		Port:     hepler.GetEnv("API_PORT"),
		LogLever: hepler.GetEnv("LOG_LEVEL"),
	}
}
