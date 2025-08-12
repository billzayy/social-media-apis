package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/billzayy/social-media/back-end/chat-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	ID             uuid.UUID
	Conn           *websocket.Conn
	Send           chan models.ReceiveSocketMsg
	ConversationId uuid.UUID
}

type WebSocketServer struct {
	Clients    map[uuid.UUID][]*Client // map[ConversationId]clients
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan models.Messages
	Lock       sync.RWMutex
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		Clients:    make(map[uuid.UUID][]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan models.Messages),
	}
}

func (ws *WebSocketServer) Run() {
	for {
		select {
		case client := <-ws.Register:
			ws.Lock.Lock()
			ws.Clients[client.ConversationId] = append(ws.Clients[client.ConversationId], client)
			ws.Lock.Unlock()
		case client := <-ws.Unregister:
			ws.removeClient(client)
		case msg := <-ws.Broadcast:
			ws.sendToRoom(msg)
		}
	}
}

func (ws *WebSocketServer) removeClient(c *Client) {
	ws.Lock.Lock()
	defer ws.Lock.Unlock()

	clients := ws.Clients[c.ConversationId]
	for i, client := range clients {
		if client == c {
			ws.Clients[c.ConversationId] = append(clients[:i], clients[i+1:]...)
			close(c.Send)
			return
		}
	}
}

func (ws *WebSocketServer) sendToRoom(msg models.Messages) {
	ws.Lock.RLock()
	defer ws.Lock.RUnlock()

	receiveMsg := models.ReceiveSocketMsg{
		SenderId:       msg.SenderId,
		ConversationId: msg.ConversationId,
		Content:        msg.Content,
		MessageType:    msg.MessageType,
		CreatedAt:      msg.CreatedAt,
	}

	clients := ws.Clients[receiveMsg.ConversationId]
	for _, client := range clients {
		if msg.ReceiverId != uuid.Nil { // Private Case
			if client.ID != msg.ReceiverId {
				continue
			}
		} else {
			if client.ID == msg.SenderId { // Group Case
				continue
			}
		}

		select {
		case client.Send <- receiveMsg:
			// fmt.Println(<-ws.Register)
		default:
			// Drop client if buffer full
			go ws.removeClient(client)
		}
	}
}

func (ws *WebSocketServer) BroadcastToRoom(ConversationId, SenderId, ReceiverId uuid.UUID, content string) error {
	msg := models.Messages{
		ConversationId: ConversationId,
		SenderId:       SenderId,
		ReceiverId:     ReceiverId,
		Content:        content,
		CreatedAt:      time.Now().String(),
	}

	ws.Broadcast <- msg

	return nil
}

func (ws *WebSocketServer) HandleWebSocket(c *gin.Context) {
	ConversationIdStr := c.Query("conversationId")
	ReceiverIdStr := c.Query("receiverId")

	ConversationId, err := uuid.Parse(ConversationIdStr)
	if err != nil {
		http.Error(c.Writer, "Invalid conversationId", http.StatusBadRequest)
		return
	}

	ReceiverId, err := uuid.Parse(ReceiverIdStr)
	if err != nil {
		http.Error(c.Writer, "Invalid senderId", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	client := &Client{
		ID:             ReceiverId,
		Conn:           conn,
		Send:           make(chan models.ReceiveSocketMsg, 256),
		ConversationId: ConversationId,
	}

	ws.Register <- client

	go client.writePump()
	go client.readPump(ws)
}

func (c *Client) readPump(ws *WebSocketServer) {
	defer func() {
		ws.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msgBytes, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		var msg models.Messages
		if err := json.Unmarshal(msgBytes, &msg); err != nil {
			continue
		}
		msg.ConversationId = c.ConversationId
		msg.SenderId = c.ID
		msg.CreatedAt = time.Now().String()
		ws.Broadcast <- msg
	}
}

func (c *Client) writePump() {
	for msg := range c.Send {
		c.Conn.WriteJSON(msg)
	}
}
