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
	t, err := opencast.Agents.Get(localConfig.Opencast.Agent)
	if err != nil {
		slog.Error("Something went wrong")
		c.JSON(http.StatusInternalServerError, nil)
		metrics.UpdateState("internal_server_error")
		return
	}

	metrics.UpdateTime()

	metrics.UpdateState(t.Status)

	c.JSON(http.StatusOK, t.Status == "capturing")
}
