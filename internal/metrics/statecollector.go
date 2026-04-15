package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	stateCollector = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "state",
		Help: "State of the CaptureAgent",
	}, []string{"state"})
)

func (col Collector) UpdateState(state string) {
	stateCollector.WithLabelValues(state).Set(1)
}
