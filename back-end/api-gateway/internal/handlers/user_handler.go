package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/billzayy/social-media/back-end/api-gateway/internal/models"
	"github.com/billzayy/social-media/back-end/api-gateway/internal/repository"
	"github.com/billzayy/social-media/back-end/user-service/api"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	port string
}

func NewUserHandler(port string) *UserHandler {
	return &UserHandler{
		port: port,
	}
}

func (uH *UserHandler) GetUserByIdHandler(c *gin.Context) {
	id := c.Query("id")

	server, client, err := repository.UserRepo(uH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := client.GetUserById(ctx, &api.ReqGetUser{Id: id})

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, resp)
}

func (uH *UserHandler) UpdateUserHandler(c *gin.Context) {
	var req models.ReqUpdateUser

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err.Error())
		return
	}

	server, client, err := repository.UserRepo(uH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	var protoBirthDate *timestamppb.Timestamp

	if !req.BirthDate.IsZero() {
		protoBirthDate = timestamppb.New(req.BirthDate)
	}

	convertedReq := &api.ReqUpdateInfo{
		Id:             req.ID,
		UserName:       req.UserName,
		FullName:       req.FullName,
		Email:          req.Email,
		ProfilePicture: req.ProfilePicture,
		Location:       req.Location,
		Website:        req.Website,
		BirthDate:      protoBirthDate,
		Theme:          req.Theme,
		Language:       req.Language,
		Country:        req.Country,
	}

	_, err = client.UpdateInfoUser(ctx, convertedReq)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, "Updated User")
}

func (uH *UserHandler) UpdatePasswordHandler(c *gin.Context) {
	var req models.ReqUpdatePassword

	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}

	server, client, err := repository.UserRepo(uH.port)

	if err != nil {
		models.Response(c, http.StatusBadRequest, "Bad Request")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	convertInput := &api.ReqUpdatePassword{
		Id:          req.ID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}

	_, err = client.UpdatePassword(ctx, convertInput)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, "Updated Successful")
}
