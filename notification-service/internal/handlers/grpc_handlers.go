package handlers

import (
	"context"
	"log"

	grpc "github.com/billzayy/social-media/notification-service/api"
	"github.com/billzayy/social-media/notification-service/internal/models"
	"github.com/billzayy/social-media/notification-service/internal/services"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcServer struct {
	grpc.UnimplementedNotificationServiceServer
	NotifyService *services.NotifyService
}

func NewNotifyGrpcServer(sv *services.Services) *GrpcServer {
	return &GrpcServer{
		NotifyService: sv.NotifyService,
	}
}

func (uG *GrpcServer) CreateSendNotify(ctx context.Context, req *grpc.ReqSendNotify) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			log.Printf("Recovered in CreateSendNotify")
		}
	}()

	senderId, err := uuid.Parse(req.SenderId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	receiverId, err := uuid.Parse(req.ReceiverId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	reqInput := models.ReqSendNotify{
		SenderId:   senderId,
		ReceiverId: receiverId,
		Messages:   req.Messages,
		Type:       models.NotificationType(req.Type),
		Url:        req.Type,
	}

	err = uG.NotifyService.SendNotifyServiceGrpc(reqInput)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	return &emptypb.Empty{}, nil
}

func (uG *GrpcServer) GetNotify(ctx context.Context, req *grpc.ReqGetNotify) (*grpc.RespNotificationsList, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in GetNotify")
		}
	}()

	dataList, err := uG.NotifyService.GetNotifyService(req.ReceiverId)

	if err != nil {
		return &grpc.RespNotificationsList{}, status.Errorf(codes.Internal, "%v", err)
	}

	resultList := &grpc.RespNotificationsList{
		Notifications: make([]*grpc.Notifications, 0),
	}

	for _, v := range dataList {
		var protoBirthDate *timestamppb.Timestamp

		if !v.Date.IsZero() {
			protoBirthDate = timestamppb.New(v.Date)
		}

		result := &grpc.Notifications{
			Id:         v.Id.String(),
			SenderId:   v.SenderId.String(),
			ReceiverId: v.ReceiverId.String(),
			Messages:   v.Messages,
			Type:       string(v.Type),
			IsRead:     v.IsRead,
			Date:       protoBirthDate,
		}

		resultList.Notifications = append(resultList.Notifications, result)
	}

	return resultList, nil
}

func (uG *GrpcServer) GetUnreadNotify(ctx context.Context, req *grpc.ReqUnreadList) (*grpc.RespGetUnreadList, error) {
	data, err := uG.NotifyService.GetUnreadNotifyService(req.UserId)

	if err != nil {
		return &grpc.RespGetUnreadList{}, status.Errorf(codes.Internal, "%v", err)
	}

	result := &grpc.RespGetUnreadList{
		Count: int64(data),
	}

	return result, nil
}

func (uG *GrpcServer) UpdateRead(ctx context.Context, req *grpc.ReqUpdateNotify) (*emptypb.Empty, error) {
	err := uG.NotifyService.UpdateReadService(req.NotifyId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%s", err)
	}

	return &emptypb.Empty{}, nil
}
