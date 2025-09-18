package handlers

import (
	"context"
	"log"

	grpc "github.com/billzayy/social-media/chat-service/api"
	"github.com/billzayy/social-media/chat-service/internal/models"
	"github.com/billzayy/social-media/chat-service/internal/services"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	grpc.UnimplementedChatServiceServer
	ChatService *services.ChatService
}

func NewChatGrpcServer(sv *services.Services) *GrpcServer {
	return &GrpcServer{
		ChatService: sv.ChatService,
	}
}
func (uG *GrpcServer) GetMessage(ctx context.Context, req *grpc.ReqGetMessage) (*grpc.RespListMessage, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			log.Printf("Recovered in GetMessage")
		}
	}()

	conversationId, err := uuid.Parse(req.ConversationId)

	if err != nil {
		return &grpc.RespListMessage{}, status.Errorf(codes.Internal, "%v", err)
	}

	data, err := uG.ChatService.GetMessage(conversationId, int(req.Limit), int(req.Offset))

	if err != nil {
		return &grpc.RespListMessage{}, status.Errorf(codes.Internal, "%v", err)
	}

	var listMessages []*grpc.RespGetMessage

	for _, v := range data {
		message := &grpc.RespGetMessage{
			Id: v.ID.String(),
			Sender: &grpc.User{
				UserId:   v.Sender.UserId.String(),
				UserName: v.Sender.UserName,
				Avatar:   v.Sender.Avatar,
			},
			Content:     v.Content,
			MessageType: v.MessageType,
			CreatedAt:   v.CreatedAt,
		}

		listMessages = append(listMessages, message)
	}

	result := &grpc.RespListMessage{
		Data: listMessages,
	}

	return result, nil
}

func (uG *GrpcServer) SendPrivateConversation(ctx context.Context, req *grpc.ReqSendPrivateMessage) (*grpc.RespSendPrivateMessage, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			log.Printf("Recovered in SendPrivateConversation")
		}
	}()

	senderId, err := uuid.Parse(req.SenderId)

	if err != nil {
		return &grpc.RespSendPrivateMessage{}, status.Errorf(codes.Internal, "%v", err)
	}

	receiverId, err := uuid.Parse(req.ReceiverId)

	if err != nil {
		return &grpc.RespSendPrivateMessage{}, status.Errorf(codes.Internal, "%v", err)
	}

	input := models.ReqSendMessage{
		SenderId:    senderId,
		ReceiverId:  receiverId,
		Content:     req.Content,
		MessageType: req.MessageType,
	}

	roomId, err := uG.ChatService.SendPrivateConversation(input)

	if err != nil {
		return &grpc.RespSendPrivateMessage{}, status.Errorf(codes.Internal, "%v", err)
	}

	return &grpc.RespSendPrivateMessage{RoomId: roomId.String()}, nil
}

func (uG *GrpcServer) DeleteMessage(ctx context.Context, req *grpc.ReqDeleteMessage) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			log.Printf("Recovered in SendPrivateConversation")
		}
	}()

	messageId, err := uuid.Parse(req.MessageId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	err = uG.ChatService.DeleteMessage(messageId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	return &emptypb.Empty{}, nil
}
