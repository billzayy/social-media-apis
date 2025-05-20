package routes

import (
	"github.com/billzayy/social-media/back-end/api-gateway/internal/handlers"
	"github.com/billzayy/social-media/back-end/api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handlers) {
	authV1 := router.Group("/auth")
	{
		authV1.POST("/login", h.AuthHandler.LoginHandler)
		authV1.POST("/register", h.AuthHandler.RegisterHandler)
		authV1.POST("/refresh-token", h.AuthHandler.RefreshTokenHandler)
	}

	router.GET("/api/v1/post/get-post", h.PostHandler.GetPostHandler)

	postV1 := router.Group("/api/v1/post", middleware.AuthMiddleware())
	{
		postV1.POST("/add-post", h.PostHandler.AddPostHandler)
		postV1.DELETE("/delete-post", h.PostHandler.DeletePostHandler)
	}

	interact := router.Group("/api/v1/interact", middleware.AuthMiddleware())
	{
		interact.POST("/check-like", h.InteractHandler.CheckLikeHandler)
		interact.POST("/add-like", h.InteractHandler.AddLikeHandler)
		interact.POST("/add-comment", h.InteractHandler.AddCommentHandler)
		interact.DELETE("/delete-like", h.InteractHandler.RemoveLikeHandler)
		interact.DELETE("/delete-comment", h.InteractHandler.DeleteCommentHandler)
	}
}
