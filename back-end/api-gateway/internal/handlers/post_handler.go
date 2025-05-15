package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/billzayy/social-media/back-end/api-gateway/internal/models"
	"github.com/billzayy/social-media/back-end/api-gateway/internal/repository"
	"github.com/billzayy/social-media/back-end/post-service/api"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostHandler struct {
	port string
}

func NewPostHandler(port string) *PostHandler {
	return &PostHandler{
		port: port,
	}
}

func (pH *PostHandler) GetPostHandler(c *gin.Context) {
	server, client, err := repository.PostRepo(pH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		fmt.Println(err)
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := client.GetPost(ctx, &emptypb.Empty{})
	if err != nil {
		fmt.Println(err)
		models.Response(c, http.StatusInternalServerError, err.Error())
		return
	}

	// fmt.Println(resp)
	models.Response(c, http.StatusOK, resp)
}

func (pH *PostHandler) AddPostHandler(c *gin.Context) {
	server, client, err := repository.PostRepo(pH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		fmt.Println(err)
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	var req api.CreatePostReq

	if err = c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err)
		return
	}

	resp, err := client.CreatePost(ctx, &req)

	models.Response(c, http.StatusOK, resp)
}

func (pH *PostHandler) DeletePostHandler(c *gin.Context) {
	server, client, err := repository.PostRepo(pH.port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	defer server.Close()

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	req := &api.DeletePostReq{
		Id: c.Query("id"),
	}

	_, err = client.DeletePost(ctx, req)

	if err != nil {
		models.Response(c, http.StatusInternalServerError, err)
		return
	}

	models.Response(c, http.StatusOK, "Delete Success")
}
