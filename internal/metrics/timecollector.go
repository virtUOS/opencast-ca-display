package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var lastUpdate time.Time

type myCollector struct {
	metric *prometheus.Desc
}

func (c *myCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.metric
}

func (c *myCollector) Collect(ch chan<- prometheus.Metric) {
	t := lastUpdate
	if t.IsZero() {
		return
	}
	s := prometheus.NewMetricWithTimestamp(t, prometheus.MustNewConstMetric(c.metric, prometheus.CounterValue, float64(t.Unix())))
	ch <- s
}

var (
	timeCollector = &myCollector{
		metric: prometheus.NewDesc(
			"last_update",
			"Timestamp of last update from CaptureAgent",
			nil,
			nil,
		),
	}
)

func UpdateTime() {
	lastUpdate = time.Now()
}
