package metrics

import "github.com/prometheus/client_golang/prometheus"

type Collector struct{}

func (col Collector) init() {
	prometheus.MustRegister(timeCollector)
	prometheus.MustRegister(timeCollector)
}
