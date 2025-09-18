package routes

import (
	"github.com/billzayy/social-media/post-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handlers) {
	postV1 := router.Group("/api/v1/post")
	{
		postV1.GET("/get-post", h.PostHandler.GetPostHandler)
		postV1.POST("/add-post", h.PostHandler.CreatePostHandler)
		postV1.DELETE("/delete-post", h.PostHandler.DeletePostHandler)
	}

	interactV1 := router.Group("/api/v1/interact")
	{
		interactV1.POST("/check-like", h.InteractHandler.CheckLikeHandler)
		interactV1.POST("/add-like", h.InteractHandler.AddLikeHandler)
		interactV1.POST("/add-comment", h.InteractHandler.AddCommentHandler)
		interactV1.DELETE("/delete-like", h.InteractHandler.RemoveLikeHandler)
		interactV1.DELETE("/delete-comment", h.InteractHandler.DeleteCommentHandler)
	}
}
