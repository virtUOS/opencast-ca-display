package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"opencast-ca-display/internal/metrics"
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
	client := &http.Client{Timeout: time.Duration(localConfig.Timeout * int(time.Millisecond))}
	url := localConfig.Opencast.URL + "/capture-admin/agents/" + localConfig.Opencast.Agent + ".json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		// stateCollector.WithLabelValues("internal_server_error").Set(1)
		metrics.UpdateState("internal_server_error")
		return
	}
	req.SetBasicAuth(localConfig.Opencast.Username, localConfig.Opencast.Password)
	resp, err := client.Do(req)
	// lastUpdate = time.Now()
	metrics.UpdateTime()
	if err != nil {
		if os.IsTimeout(err) {
			log.Println("Request timed out:", err)
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
			// stateCollector.WithLabelValues("gateway_timeout").Set(1)
			metrics.UpdateState("gateway_timeout")
		} else {
			log.Println(err)
			c.JSON(http.StatusBadGateway, gin.H{"error": "Internal server error"})
			// stateCollector.WithLabelValues("internal_server_error").Set(1)
			metrics.UpdateState("internal_server_error")
		}
		return
	}

	if resp.StatusCode != 200 {
		log.Println(resp)
		c.JSON(resp.StatusCode, nil)
		// stateCollector.WithLabelValues(fmt.Sprintf("%d", resp.StatusCode)).Set(1)
		metrics.UpdateState(fmt.Sprintf("%d", resp.StatusCode))
		return
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		// stateCollector.WithLabelValues("internal_server_error").Set(1)
		metrics.UpdateState("internal_server_error")
		return
	}
	s := string(bodyText)
	var result AgentStateResult
	jsonErr := json.Unmarshal([]byte(s), &result)

	if jsonErr != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		// stateCollector.WithLabelValues("internal_server_error").Set(1)
		metrics.UpdateState("internal_server_error")
		return
	}

	// stateCollector.Reset()
	// stateCollector.WithLabelValues(result.Update.State).Set(1)
	metrics.UpdateState(result.Update.State)

	c.JSON(http.StatusOK, result.Update.State == "capturing")
}
