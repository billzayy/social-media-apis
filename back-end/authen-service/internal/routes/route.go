package routes

import (
	"github.com/billzayy/social-media/back-end/authen-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handlers) {
	v1 := router.Group("/auth")
	{
		v1.POST("/login", h.UserHandler.LoginHandler)
		v1.POST("/register", h.UserHandler.RegisterHandler)
		v1.POST("/refresh-token", h.UserHandler.RefreshTokenHandler)
	}
}
