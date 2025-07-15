package repository

import (
	authPb "github.com/billzayy/social-media/back-end/authen-service/api"
	notificationPb "github.com/billzayy/social-media/back-end/notification-service/api"
	postPb "github.com/billzayy/social-media/back-end/post-service/api"
	userPb "github.com/billzayy/social-media/back-end/user-service/api"
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

func UserRepo(port string) (*grpc.ClientConn, userPb.UserServiceClient, error) {
	userServer, err := ConnectService("localhost", port)

	if err != nil {
		return &grpc.ClientConn{}, nil, err
	}

	client := userPb.NewUserServiceClient(userServer)

	return userServer, client, nil
}

func NotificationRepo(port string) (*grpc.ClientConn, notificationPb.NotificationServiceClient, error) {
	userServer, err := ConnectService("localhost", port)

	if err != nil {
		return &grpc.ClientConn{}, nil, err
	}

	client := notificationPb.NewNotificationServiceClient(userServer)

	return userServer, client, nil
}
