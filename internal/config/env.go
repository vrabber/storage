package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type EnvLoader struct{}

func NewEnvLoader() *EnvLoader {
	return &EnvLoader{}
}

func (e *EnvLoader) Load() (*Config, error) {
	partUploadTimeout, err := strconv.Atoi(os.Getenv("PART_UPLOAD_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse PART_UPLOAD_TIMEOUT: %w", err)
	}

	if partUploadTimeout < 0 {
		return nil, fmt.Errorf("invalid PART_UPLOAD_TIMEOUT: %d", partUploadTimeout)
	}

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
		Server: ServerConfig{
			Host:              os.Getenv("SERVER_HOST"),
			Port:              os.Getenv("SERVER_PORT"),
			PartUploadTimeout: time.Duration(partUploadTimeout) * time.Second,
		},
	}, nil
}
