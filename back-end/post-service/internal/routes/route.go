package routes

import (
	"github.com/billzayy/social-media/back-end/post-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handlers) {
	v1 := router.Group("/api/v1/post")
	{
		v1.GET("/get-post", h.PostHandler.GetPostHandler)
		v1.POST("/add-post", h.PostHandler.AddPostHandler)
		v1.DELETE("/delete-post", h.PostHandler.DeletePostHandler)
	}
}
