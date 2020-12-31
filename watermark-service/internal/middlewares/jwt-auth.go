package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watermark-services/watermark-service/internal/logger"
	"github.com/watermark-services/watermark-service/internal/service"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(bearerSchema):]

		token, err := service.ValidateToken(tokenString)

		if token.Valid {
			//claims := token.Claims.(jwt.MapClaims)
			// to Abort when user is not admin
		} else {
			logger.Log.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
