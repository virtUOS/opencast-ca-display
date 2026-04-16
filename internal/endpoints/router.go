package endpoints

import (
	"net/http"
	"opencast-ca-display/internal/config"

	"github.com/gin-gonic/gin"
)

var localConfig *config.Config

func ApiRouter(group *gin.RouterGroup) { // TODO

	localConfig = config.Default()

	group.GET("/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, localConfig.Display)
	})

	// status Endpoint
	group.GET("/status", statusEndpoint)

	// calendar Endpoint
	group.GET("/calendar", calendarEndpoint)

	// network Endpoint
	group.GET("/network_info", networkEndpoint)
}
