package routes

import (
	"github.com/billzayy/social-media/back-end/chat-service/internal/handlers"
	"github.com/billzayy/social-media/back-end/chat-service/internal/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handlers, wsServer *utils.WebSocketServer) {
	chat := router.Group("/api/v1/chat")
	{
		chat.GET("/get-message", h.ChatHandler.GetChatHandler)
		chat.POST("/send", h.ChatHandler.SendChatHandler)
		chat.DELETE("/delete", h.ChatHandler.DeleteChatHandler)
	}
	router.GET("/ws/chat", wsServer.HandleWebSocket)
}
