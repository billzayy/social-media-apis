package routes

import (
	"github.com/billzayy/social-media/back-end/notification-service/internal/handlers"
	"github.com/billzayy/social-media/back-end/notification-service/internal/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handlers) {
	notifyRoute := router.Group("/api/v1/notifications")
	{
		notifyRoute.GET("/get", h.NotifyHandler.GetNotify)
		notifyRoute.POST("/send", h.NotifyHandler.CreateNotify)
		notifyRoute.GET("/unread-count", h.NotifyHandler.GetUnreadList)  // Count unread
		notifyRoute.PATCH("/:id/read", h.NotifyHandler.UpdateReadStatus) // Read Notification
	}

	router.GET("/ws", utils.ServerWs)
}
