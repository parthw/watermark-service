package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/watermark-service/internal/service"
)

//Credentials struct
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//LoginController function
func LoginController(ctx *gin.Context) string {
	var credentials Credentials
	err := ctx.ShouldBindJSON(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := service.NewLoginService().Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return service.GenerateToken(credentials.Username, true)
	}
	return ""
}
