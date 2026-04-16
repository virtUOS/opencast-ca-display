// Opencast Capture Agent Display
// Copyright 2024 Osnabrück University
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"embed"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"opencast-ca-display/internal/config"
	"opencast-ca-display/internal/endpoints"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cConfig config.Config

	//go:embed assets
	res embed.FS
)

var (
	lastUpdate time.Time
)

type myCollector struct {
	metric *prometheus.Desc
}

func (c *myCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.metric
}

func (c *myCollector) Collect(ch chan<- prometheus.Metric) {
	t := lastUpdate
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

var (
	stateCollector = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "state",
		Help: "State of the CaptureAgent",
	}, []string{"state"})
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// disable all proxies
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// Use assets/index.html for /
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/assets/"
		r.HandleContext(c)
	})

	// Static assets
	assets, err := fs.Sub(res, "assets")
	if err != nil {
		log.Fatal(err)
	}
	r.StaticFS("/assets", http.FS(assets))

	endpoints.ApiRouter(r.Group("/"))

	return r
}

func init() {
	prometheus.MustRegister(stateCollector)
	prometheus.MustRegister(timeCollector)
}

func setupMetricsRouter() *gin.Engine {
	r := gin.Default()
	// disable all proxies
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return r
}

func main() {
	// if _, err := loadConfig("opencast-ca-display.yml"); err != nil {
	// 	log.Fatalf("Failed to load configuration: %v", err)
	// }

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	cConfig, err := cConfig.LoadFromFile("opencast-ca-display.yml")

	config.SetConfig(cConfig)
	
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if cConfig.Metrics.Enable {
		go func() {
			metricsRouter := setupMetricsRouter()
			if err := metricsRouter.Run(cConfig.Metrics.Listen); err != nil {
				log.Fatalf("Failed to run metrics server: %v", err)
			}
		}()
	}

	r := setupRouter()
	if err := r.Run(cConfig.Listen); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
