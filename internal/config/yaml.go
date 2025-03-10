package config

import (
	"fmt"
	"os"

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
	}, nil
}
