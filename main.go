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
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
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
	Info       string `json:"info"`
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
}

var (
	config Config

	//go:embed assets
	res embed.FS
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
		return nil, errors.New("No Opencast server URL in configuration")
	}

	if config.Listen == "" {
		config.Listen = "127.0.0.1:8080"
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
	assets, _ := fs.Sub(res, "assets")
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
		req.SetBasicAuth(config.Opencast.Username, config.Opencast.Password)
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadGateway, nil)
			return
		}

		if resp.StatusCode != 200 {
			log.Println(resp)
			c.JSON(resp.StatusCode, nil)
			return
		}

		bodyText, err := io.ReadAll(resp.Body)
		s := string(bodyText)
		var result AgentStateResult
		json.Unmarshal([]byte(s), &result)

		c.JSON(http.StatusOK, result.Update.State == "capturing")
	})

	return r
}

func main() {
	if _, err := loadConfig("opencast-ca-display.yml"); err != nil {
		log.Fatal(err)
		return
	}
	r := setupRouter()
	r.Run(config.Listen)
}
