package config

import (
	"errors"
)

// validate checks if the Opencast configuration has a valid URL.
// Returns an error if the URL is empty.
func (oc_conf *OpencastConfig) validate() error {
	if oc_conf.URL == "" {
		return errors.New("Opencast: No Opencast server URL in configuration")
	}
	return nil
}

// validate checks if the Display configuration is valid.
// Currently always returns nil (no validation implemented).
func (display_conf *DisplayConfig) validate() error {
	return nil
}

// validate checks if the DisplayState configuration is valid.
// Currently always returns nil (no validation implemented).
func (disp_state_conf *DisplayStateConfig) validate() error {
	return nil
}

// validate checks if the Metrics configuration is valid.
// Currently always returns nil (no validation implemented).
func (metrics_conf *MetricsConfig) validate() error {
	return nil
}

// Validate performs validation on all configuration sections.
// It checks Opencast, Display, and Metrics configurations, as well as
// the main timeout and listen settings.
// Returns a combined error of all validation failures, or nil if all are valid.
func (conf *Config) Validate() error {
	opencast_err := conf.Opencast.validate()

	display_err := conf.Display.validate()

	metrics_err := conf.Metrics.validate()

	var timeout_err error
	if conf.Timeout <= 0 {
		timeout_err = errors.New("the timeout can´t be zero or negative")
	}

	var listen_err error
	if conf.Listen == "" {
		listen_err = errors.New("it most be configured, where the main backend should listen on")
	}

	err := errors.Join(opencast_err, display_err, metrics_err, timeout_err, listen_err)

	return err
}
