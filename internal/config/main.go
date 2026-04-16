package config

import (
	"log/slog"
	"os"
	"sync/atomic"

	"gopkg.in/yaml.v3"
)

var defaultConfig atomic.Pointer[Config]

func init() {
	defaultConfig.Store(New())
}

func New() *Config {
	return &Config{}
}

// LoadFromFile loads the configuration from a YAML file at the given path.
// It validates and serializes the configuration after loading.
// Returns the loaded Config or an error if loading, parsing, validation, or serialization fails.
func (conf *Config) LoadFromFile(path string) (*Config, error) {
	yamlFile, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &conf)

	if err != nil {
		return nil, err
	}

	conf.Serialize()

	validation_err := conf.Validate()

	if validation_err != nil {
		return nil, validation_err
	}

	slog.Info("Finished loading configuration from file")

	return conf, nil
}

func SetConfig(c *Config){
	defaultConfig.Store(c)
}

func Default() *Config {
	return defaultConfig.Load()
}