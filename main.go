package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

type Config struct {
	Opencast struct {
		Url      string
		Username string
		Password string
		Agent    string
	}

	Listen string
}

var config Config

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

	// Static assets
	r.StaticFile("/", "./assets/index.html")
	r.Static("/assets", "./assets")

	// Status
	r.GET("/status", func(c *gin.Context) {
		client := &http.Client{}
		url := config.Opencast.Url + "/capture-admin/agents/" + config.Opencast.Agent + ".json"
		req, err := http.NewRequest("GET", url, nil)
		req.SetBasicAuth(config.Opencast.Username, config.Opencast.Password)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		bodyText, err := ioutil.ReadAll(resp.Body)
		s := string(bodyText)
		var result AgentStateResult
		json.Unmarshal([]byte(s), &result)

		capturing, _ := json.Marshal(result.Update.State == "capturing")
		c.String(http.StatusOK, string(capturing))
	})

	return r
}

func main() {
	if _, err := loadConfig("opencast-ca-display.yml"); err != nil {
		log.Fatal(err)
		return
	}
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(config.Listen)
}
