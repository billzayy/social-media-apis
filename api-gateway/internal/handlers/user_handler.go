package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/billzayy/social-media/api-gateway/internal/models"
	"github.com/billzayy/social-media/api-gateway/internal/repository"
	"github.com/billzayy/social-media/user-service/api"
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

// Get User By Id godoc
//
//	@Summary		Get User By Id
//	@Description	Get User with Request Id
//	@Security		BearerAuth
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"User Id"
//	@Success		200	{object}	models.SwaggerUserResp
//	@Failure		500	{object}	models.ResponseDataType
//	@Router			/api/v1/user/get-user [get]
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

	resp, err := client.GetUserById(ctx, &api.ReqGetUser{ID: id})

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, resp)
}

// Update User godoc
//
//	@Summary		Update User
//	@Description	Update User with request and token
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		models.ReqUpdateUser	true	"Request Update User"
//	@Success		200		{object}	models.ResponseDataType
//	@Failure		400		{object}	models.ResponseDataType
//	@Failure		500		{object}	models.ResponseDataType
//	@Router			/api/v1/user/update-info [patch]
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
		ID:             req.ID,
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

// Update Password godoc
//
//	@Summary		Update Password
//	@Description	Update Password with request and token
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		models.ReqUpdatePassword	true	"Request Password"
//	@Success		200		{object}	models.ResponseDataType
//	@Failure		400		{object}	models.ResponseDataType
//	@Failure		500		{object}	models.ResponseDataType
//	@Router			/api/v1/user/change-password [patch]
func (uH *UserHandler) UpdatePasswordHandler(c *gin.Context) {
	var req api.ReqUpdatePassword

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

	_, err = client.UpdatePassword(ctx, &req)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusOK, "Updated Successful")
}
