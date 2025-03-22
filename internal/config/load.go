package config

import "fmt"

const (
	SourceEnv  = "env"
	SourceYaml = "yaml"
)

var YamlConfigFile = "config.yaml"

func Load(source string) (*Config, error) {
	var loader Loader

	switch source {
	case SourceEnv:
		loader = NewEnvLoader()
	case SourceYaml:
		loader = NewYamlLoader(YamlConfigFile)
	default:
		return nil, fmt.Errorf("unsupported source %s", source)
	}

	return loader.Load()
}
