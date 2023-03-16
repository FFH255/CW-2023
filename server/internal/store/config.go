package store

import "github.com/FFH255/CW-2023/internal/hepler"

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func NewConfig() *Config {
	return &Config{
		Host:     hepler.GetEnv("DB_HOST"),
		Port:     hepler.GetEnv("DB_PORT"),
		User:     hepler.GetEnv("DB_USER"),
		Password: hepler.GetEnv("DB_PASSWORD"),
		DbName:   hepler.GetEnv("DB_NAME"),
	}
}
