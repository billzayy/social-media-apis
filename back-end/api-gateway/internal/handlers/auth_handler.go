package handlers

import (
	"context"
	"fmt"
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
		SameSite: 4,
		Secure:   true,
	})

	// fmt.Println(resp)
	models.Response(c, http.StatusOK, map[string]string{
		"UserId": resp.UserId,
		"Token":  resp.Token,
		"Type":   resp.Type,
	})
}

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
		models.Response(c, http.StatusBadRequest, err)
		return
	}

	resp, err := client.Register(ctx, &pb.RegisterReq{
		UserName:    req.UserName,
		Email:       req.Email,
		FirstName:   req.FirstName,
		SurName:     req.SurName,
		Password:    req.Password,
		Location:    req.Location,
		BirthDate:   req.BirthDate,
		Description: req.Description,
		Website:     []string{""},
	})

	if err != nil || !resp.Successful {
		fmt.Println(err)
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.Response(c, http.StatusCreated, "Account Created!")
}

func (aH *AuthHandler) RefreshTokenHandler(c *gin.Context) {
	server, client, err := repository.AuthRepo(aH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer server.Close()
	defer cancel()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	var req *pb.RefreshTokenReq

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err)
		return
	}

	resp, err := client.RefreshToken(ctx, req)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusOK, map[string]string{
		"UserId": resp.UserId,
		"Token":  resp.Token,
		"Type":   resp.Type,
	})
}
