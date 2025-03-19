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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type AgentStateResult struct {
	Update struct {
		Name  string
		State string
		Url   string
	} `json:"agent-state-update"`
}

type DisplayConfig struct {
	Text       string `json:"text"`
	Color      string `json:"color"`
	Background string `json:"background"`
	Image      string `json:"image"`
}
type Config struct {
	Opencast struct {
		Url      string
		Username string
		Password string
		Agent    string
	}

	Display struct {
		Capturing DisplayConfig `json:"capturing"`
		Idle      DisplayConfig `json:"idle"`
		Unknown   DisplayConfig `json:"unknown"`
	}

	Listen string

	Metrics struct {
		Prometheus bool
		Listen     string
	}
}

var (
	config Config

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

func loadConfig(configPath string) (*Config, error) {
	// Open config file
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// Decode YAML file
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, err
	}

	// Ensure URL does not have trailing /
	config.Opencast.Url = strings.Trim(config.Opencast.Url, "/")
	if config.Opencast.Url == "" {
		return nil, errors.New("no Opencast server URL in configuration")
	}

	if config.Listen == "" {
		config.Listen = "127.0.0.1:8080"
	}

	if config.Metrics.Listen == "" {
		config.Metrics.Listen = "0.0.0.0:9100"
	}

	return &config, nil
}

func setupRouter() *gin.Engine {
	r := gin.Default()

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

	// Display Config
	r.GET("/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, config.Display)
	})

	// Status
	r.GET("/status", func(c *gin.Context) {
		client := &http.Client{}
		url := config.Opencast.Url + "/capture-admin/agents/" + config.Opencast.Agent + ".json"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, nil)
			stateCollector.WithLabelValues("internal_server_error").Set(1)
			return
		}
		req.SetBasicAuth(config.Opencast.Username, config.Opencast.Password)
		resp, err := client.Do(req)
		lastUpdate = time.Now()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadGateway, nil)
			stateCollector.WithLabelValues("bad_gateway").Set(1)
			return
		}

		if resp.StatusCode != 200 {
			log.Println(resp)
			c.JSON(resp.StatusCode, nil)
			stateCollector.WithLabelValues(fmt.Sprintf("%d", resp.StatusCode)).Set(1)
			return
		}

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, nil)
			stateCollector.WithLabelValues("internal_server_error").Set(1)
			return
		}
		s := string(bodyText)
		var result AgentStateResult
		json_err := json.Unmarshal([]byte(s), &result)

		if json_err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, nil)
			stateCollector.WithLabelValues("internal_server_error").Set(1)
			return
		}

		stateCollector.Reset()
		stateCollector.WithLabelValues(result.Update.State).Set(1)

		c.JSON(http.StatusOK, result.Update.State == "capturing")
	})

	return r
}

func init() {
	prometheus.MustRegister(stateCollector)
	prometheus.MustRegister(timeCollector)
}

func setupMetricsRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return r
}

func main() {
	if _, err := loadConfig("opencast-ca-display.yml"); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	if config.Metrics.Prometheus {
		go func() {
			metricsRouter := setupMetricsRouter()
			if err := metricsRouter.Run(config.Metrics.Listen); err != nil {
				log.Fatalf("Failed to run metrics server: %v", err)
			}
		}()
	}

	r := setupRouter()
	if err := r.Run(config.Listen); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
