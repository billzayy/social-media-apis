package handlers

import (
	"net/http"

	"github.com/billzayy/social-media/back-end/notification-service/internal/models"
	"github.com/billzayy/social-media/back-end/notification-service/internal/services"
	"github.com/gin-gonic/gin"
)

type NotifyHandler struct {
	NotifyService *services.NotifyService
}

func NewNotifyHandler(uS *services.NotifyService) *NotifyHandler {
	return &NotifyHandler{
		NotifyService: uS,
	}
}

func (nH *NotifyHandler) CreateNotify(c *gin.Context) {
	var req models.ReqSendNotify

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadGateway, err)
		return
	}

	err := nH.NotifyService.SendNotifyService(req)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusCreated, "Created")
}

func (nH *NotifyHandler) GetNotify(c *gin.Context) {
	inputQuery := c.Query("id")

	resp, err := nH.NotifyService.GetNotifyService(inputQuery)

	if err != nil {
		models.Response(c, http.StatusNotFound, err)
		return
	}

	models.Response(c, http.StatusOK, resp)
}

func (nH *NotifyHandler) GetUnreadList(c *gin.Context) {
	input := c.Query("id")

	resp, err := nH.NotifyService.GetUnreadNotifyService(input)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	if resp == -1 {
		models.Response(c, http.StatusNotFound, "Not Found")
		return
	}

	models.Response(c, http.StatusOK, resp)
}

func (nH *NotifyHandler) UpdateReadStatus(c *gin.Context) {
	input := c.Param("id")

	err := nH.NotifyService.UpdateReadService(input)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusOK, "Updated")
}
