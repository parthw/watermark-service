package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watermark-service/internal/controller"
	"github.com/watermark-service/internal/logger"
	"github.com/watermark-service/internal/middlewares"
)

// StartServer represent starting rest-api server
func StartServer() {

	logger.Log.Info("Starting watermark service server")

	router := gin.Default()

	// No Authorization required
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "ok",
		})
	})

	// Login Endpoint: For authentication
	router.POST("/login", func(ctx *gin.Context) {
		token := controller.LoginController(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		}
	})

	// JWT Authorization required
	apiRoutes := router.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.POST("/", func(c *gin.Context) {
			controller.PutWaterMark(c)
		})
	}

	router.Run(":8080")
}
