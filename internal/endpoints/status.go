package endpoints

import (
	"log/slog"
	"net/http"
	"opencast-ca-display/internal/metrics"
	"opencast-ca-display/internal/opencast"

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
	agent, err := opencast.Agents.Get(localConfig.Opencast.Agent)
	if err != nil { // TODO enhance the error handling to distinct between errors
		slog.Error("Something went wrong")
		c.JSON(http.StatusInternalServerError, nil)
		metrics.UpdateState("internal_server_error")
		return
	}

	// update metrics to show, that the displays content was updated the last time at the current time
	metrics.UpdateTime()

	// update the state shown in the metrics endpoint
	metrics.UpdateState(agent.Status)

	// Return "true" if there is currently a recording happening and "false" otherwise
	capturing := false
	switch agent.Status {
	case "capturing":
		capturing = true
	case "uploading":
		// If the agent says, it´s uploading, there is the posibility that the agent is curretnly actually filming. 
		// This edge case is handled here.
		exists, _, err := opencast.Events.GetCurrentEvent(localConfig.Opencast.Agent)
		if err != nil {
			slog.Error("Something went wrong")
			c.JSON(http.StatusInternalServerError, nil)
			metrics.UpdateState("internal_server_error")
			return
		}
		capturing = exists
	default:
		capturing = false
	}

	c.JSON(http.StatusOK, capturing)
}
