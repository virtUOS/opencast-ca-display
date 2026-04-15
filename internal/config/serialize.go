package config

import (
	"log/slog"
	"strings"
)

// serialize cleans up the Opencast configuration by trimming trailing slashes from the URL.
func (oc_conf *OpencastConfig) serialize() {
	oc_conf.URL = strings.Trim(oc_conf.URL, "/")
}

// serialize sets default values for Metrics configuration if not specified.
// If Metrics are enabled but Listen is empty, it defaults to "0.0.0.0:9100".
func (met_conf *MetricsConfig) serialize() {
	if met_conf.Listen == "" && met_conf.Enable {
		met_conf.Listen = "0.0.0.0:9100"
		slog.Info("set metrics enpoit to default 0.0.0.0:9100")
	}
}

// serialize performs serialization on Display configuration.
// Currently a no-op (empty implementation).
func (display_conf *DisplayConfig) serialize() {}

// serialize performs serialization on DisplayState configuration.
// Currently a no-op (empty implementation).
func (display_state_conf *DisplayStateConfig) serialize() {}

// Serialize performs serialization on all configuration sections.
// It sets default values for Listen (127.0.0.1:8080) and Timeout (500) if not specified.
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
