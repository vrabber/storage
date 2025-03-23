package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type yamlConfig struct {
	Database struct {
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"database"`
	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`
	Server struct {
		Host              string `yaml:"host"`
		Port              string `yaml:"port"`
		PartUploadTimeout int    `yaml:"part_upload_timeout"`
	} `yaml:"server"`
}

type YamlLoader struct {
	filepath string
}

func NewYamlLoader(filepath string) *YamlLoader {
	return &YamlLoader{filepath: filepath}
}

func (y *YamlLoader) Load() (*Config, error) {
	data, err := os.ReadFile(y.filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var yamlCfg yamlConfig
	if err := yaml.Unmarshal(data, &yamlCfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	if yamlCfg.Server.PartUploadTimeout < 0 {
		return nil, errors.New("invalid server.part_upload_timeout")
	}

	return &Config{
		Database: DatabaseConfig{
			Name:     yamlCfg.Database.Name,
			User:     yamlCfg.Database.User,
			Password: yamlCfg.Database.Password,
			Host:     yamlCfg.Database.Host,
			Port:     yamlCfg.Database.Port,
		},
		Logging: LogConfig{
			Level: parseLogLevel(yamlCfg.Logging.Level),
		},
		Server: ServerConfig{
			Host:              yamlCfg.Server.Host,
			Port:              yamlCfg.Server.Port,
			PartUploadTimeout: time.Duration(yamlCfg.Server.PartUploadTimeout) * time.Second,
		},
	}, nil
}
