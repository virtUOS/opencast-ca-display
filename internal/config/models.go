package config

type OpencastConfig struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Agent    string `yaml:"agent"`
}

type DisplayStateConfig struct {
	Text       string `yaml:"text"`
	Color      string `yaml:"color"`
	Background string `yaml:"background"`
	Image      string `yaml:"image"`
	Info       string `yaml:"info"`
	Empty      string `yaml:"none"`
}

type DisplayConfig struct {
	Capturing DisplayStateConfig `yaml:"capturing"`
	Idle      DisplayStateConfig `yaml:"idle"`
	Unknown   DisplayStateConfig `yaml:"unknown"`
}

type MetricsConfig struct {
	Enable bool   `yaml:"enable"`
	Listen string `yaml:"listen"`
}

type Config struct {
	Opencast OpencastConfig `yaml:"opencast"`
	Display  DisplayConfig  `yaml:"display"`
	Listen   string         `yaml:"listen"`
	Timeout  int            `yaml:"timeout"`
	Metrics  MetricsConfig  `yaml:"metrics"`
}
