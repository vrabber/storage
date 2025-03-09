package config

import (
	"log/slog"
)

type DatabaseConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

type LogConfig struct {
	Level slog.Level
}

type Config struct {
	Database DatabaseConfig
	Logging  LogConfig
}

type Loader interface {
	Load() (*Config, error)
}

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		slog.Warn("invalid log level, defaulting to info", "level", level)
		return slog.LevelInfo
	}
}
