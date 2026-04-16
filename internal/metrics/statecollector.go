package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	stateCollector = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "state",
		Help: "State of the CaptureAgent",
	}, []string{"state"})
)

func UpdateState(state string) {
	// Reset all states to 0 first
	stateCollector.Reset()

	// Set the current state to 1
	stateCollector.WithLabelValues(state).Set(1)
}
