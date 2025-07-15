package utils

import (
	"net/http"

	"github.com/billzayy/social-media/back-end/api-gateway/internal/models"
	"github.com/billzayy/social-media/back-end/api-gateway/internal/repository"
	"github.com/billzayy/social-media/back-end/post-service/api"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func LikeRepoAndGetRequest(c *gin.Context, port string) (*grpc.ClientConn, api.PostServiceClient, *api.LikeRequest, error) {
	server, client, err := repository.PostRepo(port)

	if err != nil {
		return server, client, &api.LikeRequest{}, err
	}

	var req api.LikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		models.Response(c, http.StatusBadRequest, err.Error())
		return server, client, &api.LikeRequest{}, err
	}

	return server, client, &req, nil
}
