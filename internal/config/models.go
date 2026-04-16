package config

type OpencastConfig struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Agent    string `yaml:"agent"`
}

type DisplayStateConfig struct {
	Text       string `yaml:"text" json:"text"`
	Color      string `yaml:"color" json:"color"`
	Background string `yaml:"background" json:"background"`
	Image      string `yaml:"image" json:"image"`
	Info       string `yaml:"info" json:"info"`
	Empty      string `yaml:"none" json:"empty"`
}

type DisplayConfig struct {
	Capturing DisplayStateConfig `yaml:"capturing" json:"capturing"`
	Idle      DisplayStateConfig `yaml:"idle" json:"idle"`
	Unknown   DisplayStateConfig `yaml:"unknown" json:"unknown"`
}

type MetricsConfig struct {
	Enable bool   `yaml:"prometheus"` // is currently still controlled through prometheus, TODO: change in future
	Listen string `yaml:"listen"`
}

type Config struct {
	Opencast OpencastConfig `yaml:"opencast"`
	Display  DisplayConfig  `yaml:"display"`
	Listen   string         `yaml:"listen"`
	Timeout  int            `yaml:"timeout"`
	Metrics  MetricsConfig  `yaml:"metrics"`
}
