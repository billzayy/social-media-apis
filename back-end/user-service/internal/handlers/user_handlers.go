package handlers

import (
	"net/http"

	"github.com/billzayy/social-media/back-end/user-service/internal/models"
	"github.com/billzayy/social-media/back-end/user-service/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(uS *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: uS,
	}
}

func (uh *UserHandler) GetUserByIdHandler(c *gin.Context) {
	userId := c.Query("id")

	if userId == "" {
		models.Response(c, http.StatusBadRequest, "id not found")
		return
	}

	data, err := uh.UserService.GetUserService(userId)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, data)
}

func (uh *UserHandler) UpdateUserHandler(c *gin.Context) {
	var req models.ReqUpdateUser

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err.Error())
		return
	}

	err := uh.UserService.UpdateUserService(req)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, "Update Successful")
}

func (uh *UserHandler) UpdatePasswordHandler(c *gin.Context) {
	var req models.ReqUpdatePassword

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err.Error())
		return
	}

	err := uh.UserService.UpdatePasswordService(req)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, "Update Successful")
}
