package handlers

import (
	"context"
	"log"

	grpc "github.com/billzayy/social-media/back-end/user-service/api"
	"github.com/billzayy/social-media/back-end/user-service/internal/models"
	"github.com/billzayy/social-media/back-end/user-service/internal/services"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcServer struct {
	grpc.UnimplementedUserServiceServer
	UserService *services.UserService
}

func NewUserGrpcServer(sv *services.Services) *GrpcServer {
	return &GrpcServer{
		UserService: sv.UserService,
	}
}

func (uG *GrpcServer) GetUserById(ctx context.Context, req *grpc.ReqGetUser) (*grpc.RespGetUser, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in UpdateInfoUser")
		}
	}()

	data, err := uG.UserService.GetUserService(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to Get User : %v", err)
	}

	var protoBirthDate *timestamppb.Timestamp

	if !data.BirthDate.IsZero() {
		protoBirthDate = timestamppb.New(data.BirthDate)
	}

	result := &grpc.RespGetUser{
		Id:             data.ID.String(),
		UserName:       data.UserName,
		FullName:       data.FullName,
		Email:          data.Email,
		UserSince:      data.UserSince.String(),
		LastLogin:      data.LastLogin.String(),
		Active:         data.Active,
		ProfilePicture: data.ProfilePicture,
		Description:    data.Description,
		Location:       data.Location,
		Website:        data.Website,
		BirthDate:      protoBirthDate,
		Theme:          data.Theme,
		Language:       data.Language,
		Country:        data.Country,
	}

	return result, nil
}

func (uG *GrpcServer) UpdateInfoUser(ctx context.Context, req *grpc.ReqUpdateInfo) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in UpdateInfoUser")
		}
	}()

	input := models.ReqUpdateUser{
		ID:             uuid.MustParse(req.Id),
		UserName:       req.UserName,
		FullName:       req.FullName,
		Email:          req.Email,
		ProfilePicture: req.ProfilePicture,
		Description:    req.Description,
		Location:       req.Location,
		Website:        req.Website,
		BirthDate:      req.BirthDate.AsTime(),
		Theme:          req.Theme,
		Language:       req.Language,
		Country:        req.Country,
	}

	err := uG.UserService.UpdateUserService(input)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	return &emptypb.Empty{}, nil
}

func (uG *GrpcServer) UpdatePassword(ctx context.Context, req *grpc.ReqUpdatePassword) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in UpdateInfoUser")
		}
	}()

	input := models.ReqUpdatePassword{
		ID:          uuid.MustParse(req.Id),
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}

	err := uG.UserService.UpdatePasswordService(input)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	return &emptypb.Empty{}, nil
}
