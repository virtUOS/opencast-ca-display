package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

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

	return conf, nil
}
