package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/billzayy/social-media/back-end/api-gateway/internal/models"
	"github.com/billzayy/social-media/back-end/api-gateway/internal/repository"
	"github.com/billzayy/social-media/back-end/api-gateway/internal/utils"
	"github.com/billzayy/social-media/back-end/notification-service/api"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	port string
}

func NewNotificationHandler(port string) *NotificationHandler {
	return &NotificationHandler{
		port: port,
	}
}

// Send Notifications doc
//
//	@Summary		Send Notification
//	@Description	Send Notification w body data
//	@Tags			notification
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		models.SwaggerReqSendNotify	true	"Send Notify Request"
//	@Success		200		{object}	models.ResponseDataType
//	@Failure		400		{object}	models.ResponseDataType
//	@Failure		500		{object}	models.ResponseDataType
//	@Router			/api/v1/notifications/send [post]
func (nH *NotificationHandler) SendNotify(c *gin.Context) {
	server, client, err := repository.NotificationRepo(nH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}
	var req api.ReqSendNotify

	if err = c.ShouldBind(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err)
		return
	}

	_, err = client.CreateSendNotify(ctx, &api.ReqSendNotify{
		SenderId:   req.SenderId,
		ReceiverId: req.ReceiverId,
		Messages:   req.Messages,
		Type:       req.Type,
		Url:        req.Url,
	})

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	//* Send to websocket
	go utils.SendNotifyToWebSocket(req.ReceiverId, req.Messages)

	models.Response(c, http.StatusCreated, "Sent")
}

// Get Notifications doc
//
//	@Summary		Get Notification
//	@Description	Get Notification List w params
//	@Tags			notification
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	query		string	true	"UserId Request"
//	@Success		200	{object}	models.ResponseDataType
//	@Failure		400	{object}	models.ResponseDataType
//	@Failure		500	{object}	models.ResponseDataType
//	@Router			/api/v1/notifications/get [get]
func (nH *NotificationHandler) GetNotify(c *gin.Context) {
	server, client, err := repository.NotificationRepo(nH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	id := c.Query("id")

	if len(id) == 0 {
		models.Response(c, http.StatusBadRequest, "Bad Request")
		return
	}

	req := &api.ReqGetNotify{
		ReceiverId: id,
	}

	data, err := client.GetNotify(ctx, req)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusOK, &data)
}

// Count Unread Notify doc
//
//	@Summary		Count Unread Notification
//	@Description	Count Unread Notification w Id
//	@Tags			notification
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	query		string	true	"UserId Request"
//	@Success		200	{object}	models.ResponseDataType
//	@Failure		400	{object}	models.ResponseDataType
//	@Failure		500	{object}	models.ResponseDataType
//	@Router			/api/v1/notifications/unread-count [get]
func (nH *NotificationHandler) GetUnreadNotify(c *gin.Context) {
	server, client, err := repository.NotificationRepo(nH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	id := c.Query("id")

	if len(id) == 0 {
		models.Response(c, http.StatusBadRequest, "Bad Request")
		return
	}

	req := &api.ReqUnreadList{
		UserId: id,
	}

	data, err := client.GetUnreadNotify(ctx, req)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusOK, data)
}

//	 Update Read Status Notification doc
//
//	@Summary		Update Read Status Notification
//	@Description	Update Read Status Notification w param Id
//	@Tags			notification
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"NotifyId Request"
//	@Success		200	{object}	models.ResponseDataType
//	@Failure		400	{object}	models.ResponseDataType
//	@Failure		500	{object}	models.ResponseDataType
//	@Router			/api/v1/notifications/{id}/read [patch]
func (nH *NotificationHandler) UpdateRead(c *gin.Context) {
	server, client, err := repository.NotificationRepo(nH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	id := c.Param("id")

	if len(id) == 0 {
		models.Response(c, http.StatusBadRequest, "Bad Request")
		return
	}

	req := &api.ReqUpdateNotify{
		NotifyId: id,
	}

	_, err = client.UpdateRead(ctx, req)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusOK, "Updated")
}
