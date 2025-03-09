package config

import (
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		level    string
		expected slog.Level
	}{
		{
			name:     "debug_level",
			level:    "debug",
			expected: slog.LevelDebug,
		},
		{
			name:     "info_level",
			level:    "info",
			expected: slog.LevelInfo,
		},
		{
			name:     "warn_level",
			level:    "warn",
			expected: slog.LevelWarn,
		},
		{
			name:     "error_level",
			level:    "error",
			expected: slog.LevelError,
		},
		{
			name:     "empty_level",
			level:    "",
			expected: slog.LevelInfo,
		},
		{
			name:     "invalid_level",
			level:    "invalid",
			expected: slog.LevelInfo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseLogLevel(tt.level)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestEnvConfigLoader(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected *Config
	}{
		{
			name: "all_env_vars_set_debug",
			envVars: map[string]string{
				"DB_NAME":     "testdb",
				"DB_USER":     "testuser",
				"DB_PASSWORD": "testpass",
				"DB_HOST":     "testhost",
				"DB_PORT":     "5432",
				"LOG_LEVEL":   "debug",
			},
			expected: &Config{
				Database: DatabaseConfig{
					Name:     "testdb",
					User:     "testuser",
					Password: "testpass",
					Host:     "testhost",
					Port:     "5432",
				},
				Logging: LogConfig{
					Level: slog.LevelDebug,
				},
			},
		},
		{
			name: "all_env_vars_set_error",
			envVars: map[string]string{
				"DB_NAME":     "testdb",
				"DB_USER":     "testuser",
				"DB_PASSWORD": "testpass",
				"DB_HOST":     "testhost",
				"DB_PORT":     "5432",
				"LOG_LEVEL":   "error",
			},
			expected: &Config{
				Database: DatabaseConfig{
					Name:     "testdb",
					User:     "testuser",
					Password: "testpass",
					Host:     "testhost",
					Port:     "5432",
				},
				Logging: LogConfig{
					Level: slog.LevelError,
				},
			},
		},
		{
			name: "missing_log_level_defaults_to_info",
			envVars: map[string]string{
				"DB_NAME":     "testdb",
				"DB_USER":     "testuser",
				"DB_PASSWORD": "testpass",
				"DB_HOST":     "testhost",
				"DB_PORT":     "5432",
			},
			expected: &Config{
				Database: DatabaseConfig{
					Name:     "testdb",
					User:     "testuser",
					Password: "testpass",
					Host:     "testhost",
					Port:     "5432",
				},
				Logging: LogConfig{
					Level: slog.LevelInfo,
				},
			},
		},
		{
			name: "empty_env_vars",
			envVars: map[string]string{
				"DB_NAME":     "",
				"DB_USER":     "",
				"DB_PASSWORD": "",
				"DB_HOST":     "",
				"DB_PORT":     "",
				"LOG_LEVEL":   "",
			},
			expected: &Config{
				Logging: LogConfig{
					Level: slog.LevelInfo,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envVars {
				t.Setenv(k, v)
			}

			loader := NewEnvLoader()
			cfg, err := loader.Load()

			require.NoError(t, err)
			assert.Equal(t, tt.expected, cfg)
		})
	}
}

func TestYamlConfigLoader(t *testing.T) {
	tests := []struct {
		name        string
		yamlContent string
		expectErr   bool
		expected    *Config
	}{
		{
			name: "valid_config_debug",
			yamlContent: `
database:
  name: "testdb"
  user: "testuser"
  password: "testpass"
  host: "testhost"
  port: "5432"
logging:
  level: "debug"
`,
			expectErr: false,
			expected: &Config{
				Database: DatabaseConfig{
					Name:     "testdb",
					User:     "testuser",
					Password: "testpass",
					Host:     "testhost",
					Port:     "5432",
				},
				Logging: LogConfig{
					Level: slog.LevelDebug,
				},
			},
		},
		{
			name: "valid_config_error",
			yamlContent: `
database:
  name: "testdb"
  user: "testuser"
  password: "testpass"
  host: "testhost"
  port: "5432"
logging:
  level: "error"
`,
			expectErr: false,
			expected: &Config{
				Database: DatabaseConfig{
					Name:     "testdb",
					User:     "testuser",
					Password: "testpass",
					Host:     "testhost",
					Port:     "5432",
				},
				Logging: LogConfig{
					Level: slog.LevelError,
				},
			},
		},
		{
			name: "missing_log_level",
			yamlContent: `
database:
  name: "testdb"
  user: "testuser"
  password: "testpass"
  host: "testhost"
  port: "5432"
`,
			expectErr: false,
			expected: &Config{
				Database: DatabaseConfig{
					Name:     "testdb",
					User:     "testuser",
					Password: "testpass",
					Host:     "testhost",
					Port:     "5432",
				},
				Logging: LogConfig{
					Level: slog.LevelInfo,
				},
			},
		},
		{
			name: "invalid_yaml",
			yamlContent: `
invalid yaml content
---
`,
			expectErr: true,
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			tmpFile := filepath.Join(tmpDir, "config.yaml")
			err := os.WriteFile(tmpFile, []byte(tt.yamlContent), 0644)
			require.NoError(t, err)

			loader := NewYamlLoader(tmpFile)
			cfg, err := loader.Load()

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, cfg)
		})
	}
}

func TestYamlConfigLoader_FileNotFound(t *testing.T) {
	loader := NewYamlLoader("nonexistent.yaml")
	_, err := loader.Load()
	assert.Error(t, err)
}
