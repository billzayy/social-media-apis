package routes

import (
	"github.com/billzayy/social-media/authen-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handlers) {
	v1 := router.Group("/auth")
	{
		v1.POST("/login", h.AuthHandler.LoginHandler)
		v1.POST("/register", h.AuthHandler.RegisterHandler)
		v1.POST("/refresh-token", h.AuthHandler.RefreshTokenHandler)
	}
}
