package metrics

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// disable all proxies
	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return router
}
