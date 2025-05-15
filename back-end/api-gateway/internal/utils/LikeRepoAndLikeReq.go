package utils

import (
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

	var req models.LikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		return server, client, &api.LikeRequest{}, err
	}

	input := &api.LikeRequest{
		UserId: req.UserId,
		PostId: req.PostId,
	}
	return server, client, input, nil
}
