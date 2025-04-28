package handlers

import (
	"context"
	"fmt"
	"log"
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
	server, client, err := repository.AuthRepo(aH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		fmt.Println(err)
		models.ResponseUser(c, http.StatusInternalServerError, err)
		return
	}

	resp, err := client.Login(ctx, &pb.LoginReq{UserName: c.Query("userName"), Password: c.Query("password")})
	if err != nil {
		fmt.Println(err)
		models.ResponseUser(c, http.StatusInternalServerError, err)
		return
	}

	// fmt.Println(resp)
	models.ResponseUser(c, http.StatusOK, resp)
}

func (aH *AuthHandler) RegisterHandler(c *gin.Context) {
	server, client, err := repository.AuthRepo(aH.port)

	if err != nil {
		log.Fatalf("%v", err)
	}

	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var req pb.RegisterReq

	if err := c.ShouldBindJSON(&req); err != nil {
		models.ResponseUser(c, http.StatusBadRequest, err)
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

	if err != nil {
		fmt.Println(err)
		models.ResponseUser(c, http.StatusInternalServerError, err.Error())
		return
	}

	models.ResponseUser(c, http.StatusCreated, resp)
}
