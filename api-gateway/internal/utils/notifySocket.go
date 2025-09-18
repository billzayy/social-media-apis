package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[string]*websocket.Conn) // userID -> connection

func NotifySocketServer(c *gin.Context) {
	userID := c.Query("userId")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	clients[userID] = conn

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("User %s disconnected", userID)
			delete(clients, userID)
			break
		}
	}
}

func SendNotifyToWebSocket(userID string, message string) {
	if conn, ok := clients[userID]; ok {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Write error:", err)
			conn.Close()
			delete(clients, userID)
		}
	}
}
