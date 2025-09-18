package repository

import (
	authPb "github.com/billzayy/social-media/authen-service/api"
	chatPb "github.com/billzayy/social-media/chat-service/api"
	notificationPb "github.com/billzayy/social-media/notification-service/api"
	postPb "github.com/billzayy/social-media/post-service/api"
	userPb "github.com/billzayy/social-media/user-service/api"
	"google.golang.org/grpc"
)

func AuthRepo(port string) (*grpc.ClientConn, authPb.AuthServiceClient, error) {
	authServer, err := ConnectService("authen-service", port) // localhost if want to test

	if err != nil {
		return &grpc.ClientConn{}, nil, err
	}

	client := authPb.NewAuthServiceClient(authServer)

	return authServer, client, nil
}

func PostRepo(port string) (*grpc.ClientConn, postPb.PostServiceClient, error) {
	postServer, err := ConnectService("post-service", port)

	if err != nil {
		return &grpc.ClientConn{}, nil, err
	}

	client := postPb.NewPostServiceClient(postServer)

	return postServer, client, nil
}

func UserRepo(port string) (*grpc.ClientConn, userPb.UserServiceClient, error) {
	userServer, err := ConnectService("user-service", port)

	if err != nil {
		return &grpc.ClientConn{}, nil, err
	}

	client := userPb.NewUserServiceClient(userServer)

	return userServer, client, nil
}

func NotificationRepo(port string) (*grpc.ClientConn, notificationPb.NotificationServiceClient, error) {
	userServer, err := ConnectService("notification-service", port)

	if err != nil {
		return &grpc.ClientConn{}, nil, err
	}

	client := notificationPb.NewNotificationServiceClient(userServer)

	return userServer, client, nil
}

func ChatRepo(port string) (*grpc.ClientConn, chatPb.ChatServiceClient, error) {
	chatServer, err := ConnectService("chat-service", port)

	if err != nil {
		return &grpc.ClientConn{}, nil, err
	}

	client := chatPb.NewChatServiceClient(chatServer)

	return chatServer, client, nil
}
