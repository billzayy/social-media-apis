package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/billzayy/social-media/back-end/api-gateway/internal/models"
	"github.com/billzayy/social-media/back-end/api-gateway/internal/repository"
	"github.com/billzayy/social-media/back-end/chat-service/api"
	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	port string
}

func NewChatHandler(port string) *ChatHandler {
	return &ChatHandler{
		port: port,
	}
}

// GetMessage godoc
//
//	@Summary		Get Message
//	@Description	Get Message w request
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			roomId	query		string	true	"Room Id"
//	@Param			limit	query		string	true	"Limit Get"
//	@Param			offset	query		string	true	"Offset for pagination"
//	@Success		200		{object}	models.SwaggerRespGetMessage
//	@Failure		500		{object}	models.ResponseDataType
//	@Router			/api/v1/chat/get-message [get]
func (cH *ChatHandler) GetMessageHandler(c *gin.Context) {
	server, client, err := repository.ChatRepo(cH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	roomId := c.Query("roomId")
	limitQuery := c.Query("limit")
	offsetQuery := c.Query("offset")

	if len(roomId) == 0 {
		models.Response(c, http.StatusBadRequest, "Bad Request")
		return
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		return
	}

	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		return
	}

	req := &api.ReqGetMessage{
		ConversationId: roomId,
		Limit:          int64(limit),
		Offset:         int64(offset),
	}

	data, err := client.GetMessage(ctx, req)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusOK, data.GetData())
}

// SendMessage godoc
//
//	@Summary		Send Message
//	@Description	Send Message w request
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			chatType	query		string							true	"Type of room chat"
//	@Param			request		body		models.SwaggerReqSendMessage	true	"Send Chat Request"
//	@Success		200			{object}	models.ResponseDataType
//	@Failure		500			{object}	models.ResponseDataType
//	@Router			/api/v1/chat/send [post]
func (cH *ChatHandler) SendMessageHandler(c *gin.Context) {
	server, client, err := repository.ChatRepo(cH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}
	status := c.Query("chatType")

	var request *api.ReqSendPrivateMessage

	if err := c.ShouldBindJSON(&request); err != nil {
		models.Response(c, http.StatusBadRequest, err.Error())
		return
	}

	if status == "private" {
		resp, err := client.SendPrivateConversation(ctx, request)

		if err != nil {
			models.Response(c, http.StatusInternalServerError, err.Error())
			return
		}

		models.Response(c, http.StatusOK, resp)
	}
}

// DeleteMessage godoc
//
//	@Summary		Delete Message
//	@Description	Delete Message w request
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			messageId	query		string	true	"Message Id"
//	@Success		200			{object}	models.ResponseDataType
//	@Failure		500			{object}	models.ResponseDataType
//	@Router			/api/v1/chat/delete [delete]
func (cH *ChatHandler) DeleteMessageHandler(c *gin.Context) {
	server, client, err := repository.ChatRepo(cH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	id := c.Query("messageId")

	if id == "" {
		models.Response(c, http.StatusBadRequest, fmt.Errorf("request empty"))
		return
	}

	_, err = client.DeleteMessage(ctx, &api.ReqDeleteMessage{MessageId: id})

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, "Deleted!")
}
