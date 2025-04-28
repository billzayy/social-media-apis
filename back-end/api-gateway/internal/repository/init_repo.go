package repository

import (
	authPb "github.com/billzayy/social-media/back-end/authen-service/api"
	postPb "github.com/billzayy/social-media/back-end/post-service/api"
	"google.golang.org/grpc"
)

func AuthRepo(port string) (*grpc.ClientConn, authPb.AuthServiceClient, error) {
	authServer, err := ConnectService("localhost", port)

	if err != nil {
		return &grpc.ClientConn{}, nil, err
	}

	client := authPb.NewAuthServiceClient(authServer)

	return authServer, client, nil
}

func PostRepo(port string) (*grpc.ClientConn, postPb.PostServiceClient, error) {
	postServer, err := ConnectService("localhost", port)

	if err != nil {
		return &grpc.ClientConn{}, nil, err
	}

	client := postPb.NewPostServiceClient(postServer)

	return postServer, client, nil
}
