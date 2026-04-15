package endpoints

import "github.com/gin-gonic/gin"

func SetupRouter(group gin.RouterGroup) {
	// status Endpoint
	group.GET("/status", statusEndpoint)

	// calendar Endpoint
	group.GET("/calendar", calendarEndpoint)

	// network Endpoint
	group.GET("/network_info", networkEndpoint)
}
