package endpoints

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type AgentStateResult struct {
	Update struct {
		Name  string
		State string
		Url   string
	} `json:"agent-state-update"`
}

func statusEndpoint(c *gin.Context) {
	client := &http.Client{Timeout: time.Duration(config.Timeout * int(time.Millisecond))}
	url := config.Opencast.Url + "/capture-admin/agents/" + config.Opencast.Agent + ".json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		// stateCollector.WithLabelValues("internal_server_error").Set(1)
		return
	}
	req.SetBasicAuth(config.Opencast.Username, config.Opencast.Password)
	resp, err := client.Do(req)
	// lastUpdate = time.Now()
	if err != nil {
		if os.IsTimeout(err) {
			log.Println("Request timed out:", err)
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
			// stateCollector.WithLabelValues("gateway_timeout").Set(1)
		} else {
			log.Println(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": "Internal server error"})
			// stateCollector.WithLabelValues("internal_server_error").Set(1)
		}
		return
	}

	if resp.StatusCode != 200 {
		log.Println(resp)
		c.JSON(resp.StatusCode, nil)
		// stateCollector.WithLabelValues(fmt.Sprintf("%d", resp.StatusCode)).Set(1)
		return
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		// stateCollector.WithLabelValues("internal_server_error").Set(1)
		return
	}
	s := string(bodyText)
	var result AgentStateResult
	jsonErr := json.Unmarshal([]byte(s), &result)

	if jsonErr != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		// stateCollector.WithLabelValues("internal_server_error").Set(1)
		return
	}

	// stateCollector.Reset()
	// stateCollector.WithLabelValues(result.Update.State).Set(1)

	c.JSON(http.StatusOK, result.Update.State == "capturing")
}
