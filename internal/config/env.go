package config

import "os"

type EnvLoader struct{}

func NewEnvLoader() *EnvLoader {
	return &EnvLoader{}
}

func (e *EnvLoader) Load() (*Config, error) {
	return &Config{
		Database: DatabaseConfig{
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
		Logging: LogConfig{
			Level: parseLogLevel(os.Getenv("LOG_LEVEL")),
		},
	}, nil
}
