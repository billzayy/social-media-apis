package routes

import (
	"github.com/billzayy/social-media/back-end/user-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handlers) {
	userRoute := router.Group("/api/user")
	{
		userRoute.GET("/get-user", h.UserHandler.GetUserByIdHandler)
		userRoute.PATCH("/update", h.UserHandler.UpdateUserHandler)
		userRoute.PATCH("/change-password", h.UserHandler.UpdatePasswordHandler)
	}
}
