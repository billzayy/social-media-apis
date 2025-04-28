package routes

import (
	"github.com/billzayy/social-media/back-end/api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handlers) {
	authV1 := router.Group("/auth")
	{
		authV1.POST("/login", h.AuthHandler.LoginHandler)
		authV1.POST("/register", h.AuthHandler.RegisterHandler)
	}

	postV1 := router.Group("/api/v1/post")
	{
		postV1.GET("/get-post", h.PostHandler.GetPostHandler)
		postV1.POST("/add-post", h.PostHandler.AddPostHandler)
	}
}
