package config

import (
	"log/slog"
	"strings"
)

func (oc_conf *OpencastConfig) serialize() {
	oc_conf.URL = strings.Trim(oc_conf.URL, "/")
}

func (met_conf *MetricsConfig) serialize() {
	if met_conf.Listen == "" && met_conf.Enable {
		met_conf.Listen = "0.0.0.0:9100"
		slog.Info("set metrics enpoit to default 0.0.0.0:9100")
	}
}

func (display_conf *DisplayConfig) serialize() {}

func (display_state_conf *DisplayStateConfig) serialize() {}

func (conf *Config) Serialize() {
	// Serialize opencast config
	conf.Opencast.serialize()

	// Serilaize Display config
	conf.Display.serialize()

	// Serialize metrics config
	conf.Metrics.serialize()

	if conf.Listen == "" {
		conf.Listen = "127.0.0.1:8080"
		slog.Info("set backend listen to default 127.0.0.1:8080")
	}

	if conf.Timeout <= 0 {
		conf.Timeout = 500
		slog.Info("set request timeout to a reasonable level")
	}
}
