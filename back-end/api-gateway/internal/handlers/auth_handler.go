package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/billzayy/social-media/back-end/api-gateway/internal/models"
	"github.com/billzayy/social-media/back-end/api-gateway/internal/repository"
	pb "github.com/billzayy/social-media/back-end/authen-service/api"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	port string
}

func NewAuthHandler(port string) *AuthHandler {
	return &AuthHandler{
		port: port,
	}
}

// Login User godoc
//
//	@Summary		Login User
//	@Description	Login with username and password
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			userName	query		string	true	"User Name"
//	@Param			password	query		string	true	"Password"
//	@Success		200			{object}	models.SwaggerLoginResp
//	@Failure		500			{object}	models.ResponseDataType
//	@Router			/auth/login [post]
func (aH *AuthHandler) LoginHandler(c *gin.Context) {
	userName := c.Query("userName")
	password := c.Query("password")

	server, client, err := repository.AuthRepo(aH.port) // Init Repository

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := client.Login(ctx, &pb.LoginReq{
		UserName: userName,
		Password: password,
	}) // Call Login function from GRPC

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    resp.Cookie,
		Expires:  time.Now().Add(168 * time.Hour),
		HttpOnly: true, // If set true, the Front-End Js can not get the cookie
		Path:     "/",
		SameSite: 3,
		Secure:   true,
	})

	result := &pb.LoginResp{
		Token:   resp.Token,
		Type:    resp.Type,
		Expires: resp.Expires,
		User: &pb.UserResp{
			ID:             resp.User.ID,
			FullName:       resp.User.FullName,
			Email:          resp.User.Email,
			ProfilePicture: resp.User.ProfilePicture,
		},
	}

	models.Response(c, http.StatusOK, result)
}

// Register Account godoc
//
//	@Summary		Register Account
//	@Description	Register User Account
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.SwaggerLoginReq	true	"Register Request"
//	@Success		200		{object}	models.ResponseDataType
//	@Failure		400		{object}	models.ResponseDataType
//	@Failure		500		{object}	models.ResponseDataType
//	@Router			/auth/register [post]
func (aH *AuthHandler) RegisterHandler(c *gin.Context) {
	server, client, err := repository.AuthRepo(aH.port)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var req pb.RegisterReq

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := client.Register(ctx, &req)

	if err != nil || !resp.Successful {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusCreated, "Account Created!")
}

// Refresh Token godoc
//
//	@Summary		Refresh Token
//	@Description	Refresh Token API
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.ResponseDataType
//	@Failure		400	{object}	models.ResponseDataType
//	@Failure		500	{object}	models.ResponseDataType
//	@Router			/auth/refresh-token [post]
func (aH *AuthHandler) RefreshTokenHandler(c *gin.Context) {
	server, client, err := repository.AuthRepo(aH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer server.Close()
	defer cancel()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	getToken, err := c.Cookie("jwt")

	if err != nil {
		models.Response(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := client.RefreshToken(ctx, &pb.RefreshTokenReq{RefreshToken: getToken})

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    resp.Cookie,
		Expires:  time.Now().Add(168 * time.Hour),
		HttpOnly: true, // If set true, the Front-End Js can not get the cookie
		Path:     "/",
		SameSite: 3,
		Secure:   true,
	})

	models.Response(c, http.StatusOK, pb.RefreshTokenResp{
		User: &pb.UserResp{
			ID:             resp.User.ID,
			FullName:       resp.User.FullName,
			Email:          resp.User.Email,
			ProfilePicture: resp.User.ProfilePicture,
		},
		Token:   resp.Token,
		Type:    resp.Type,
		Expires: resp.Expires,
	})
}
