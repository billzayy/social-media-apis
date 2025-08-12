package handlers

import (
	"net/http"
	"strconv"

	"github.com/billzayy/social-media/back-end/chat-service/internal/models"
	"github.com/billzayy/social-media/back-end/chat-service/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChatHandler struct {
	ChatService *services.ChatService
}

type RespSent struct {
	ID     uuid.UUID `json:"roomId"`
	Status string    `json:"status"`
}

func NewChatHandler(cS *services.ChatService) *ChatHandler {
	return &ChatHandler{
		ChatService: cS,
	}
}

func (cH *ChatHandler) SendChatHandler(c *gin.Context) {
	chatType := c.Query("chatType")
	var input models.ReqSendMessage

	if err := c.ShouldBindJSON(&input); err != nil {
		models.Response(c, http.StatusBadRequest, err)
		return
	}

	if chatType == "private" {
		roomId, err := cH.ChatService.SendPrivateConversation(input)

		if err != nil || roomId == uuid.Nil {
			models.Response(c, http.StatusInternalServerError, err.Error())
			return
		}

		models.Response(c, http.StatusCreated, RespSent{
			ID:     roomId,
			Status: "Sent Successful",
		})
	}
}

func (cH *ChatHandler) GetChatHandler(c *gin.Context) {
	conversationId := c.Query("roomId")
	limitReq := c.Query("limit")
	offsetReq := c.Query("offset")

	reqId, err := uuid.Parse(conversationId)

	if err != nil {
		models.Response(c, http.StatusBadRequest, err)
		return
	}

	if limitReq == "" || offsetReq == "" {
		models.Response(c, http.StatusBadRequest, "Limit or Offset is empty")
		return
	}

	limit, err := strconv.Atoi(limitReq)

	if err != nil {
		models.Response(c, http.StatusBadRequest, "error on convert Limit")
		return
	}

	offset, err := strconv.Atoi(offsetReq)

	if err != nil {
		models.Response(c, http.StatusBadRequest, "error on convert Offset")
		return
	}

	data, err := cH.ChatService.GetMessage(reqId, limit, offset)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, data)
}

func (cH *ChatHandler) DeleteChatHandler(c *gin.Context) {
	id := c.Query("messageId")

	if id == "" {
		models.Response(c, http.StatusBadRequest, "")
		return
	}

	messageId, err := uuid.Parse(id)

	if err != nil {
		models.Response(c, http.StatusBadRequest, err.Error())
		return
	}

	err = cH.ChatService.DeleteMessage(messageId)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, "Deleted!")
}
