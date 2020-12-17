package server

import (
	"github.com/gin-gonic/gin"
	"github.com/watermark-services/watermark-service/internal/controller"
	"github.com/watermark-services/watermark-service/internal/logger"
)

// StartServer represent starting rest-api server
func StartServer() {
	logger.Log.Info("Starting gin server")
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "ok",
		})
	})

	router.POST("/", func(c *gin.Context) {
		controller.PutWaterMark(c)
	})

	router.Run(":8080")
}
