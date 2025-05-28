package handlers

import (
	"time"

	grpc "github.com/billzayy/social-media/back-end/user-service/api"
	"github.com/billzayy/social-media/back-end/user-service/internal/models"
	"github.com/billzayy/social-media/back-end/user-service/internal/services"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	grpc.UnimplementedUserServiceServer
	UserService *services.UserService
}

func NewPostGrpcServer(sv *services.Services) *GrpcServer {
	return &GrpcServer{
		UserService: sv.UserService,
	}
}

func (uG *GrpcServer) GetUserById(req *grpc.ReqGetUser) (*grpc.RespGetUser, error) {
	data, err := uG.UserService.GetUserService(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to Get User : %v", err)
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
		BirthDate:      data.BirthDate.String(),
		Theme:          data.Theme,
		Language:       data.Language,
		Country:        data.Country,
	}

	return result, nil
}

func (uG *GrpcServer) UpdateInfoUser(req *grpc.ReqUpdateInfo) error {
	birth, err := time.Parse(time.RFC3339, req.BirthDate)

	if err != nil {
		return status.Errorf(codes.Internal, "failed to convert time : %v", err)
	}

	input := models.ReqUpdateUser{
		ID:             uuid.MustParse(req.Id),
		UserName:       req.UserName,
		FullName:       req.FullName,
		Email:          req.Email,
		ProfilePicture: req.ProfilePicture,
		Description:    req.Description,
		Location:       req.Location,
		Website:        req.Website,
		BirthDate:      birth,
		Theme:          req.Theme,
		Language:       req.Language,
		Country:        req.Country,
	}

	err = uG.UserService.UpdateUserService(input)

	if err != nil {
		return status.Errorf(codes.Internal, "failed to Update User : %v", err)
	}

	return nil
}

func (uG *GrpcServer) UpdatePassword(req *grpc.ReqUpdatePassword) error {
	input := models.ReqUpdatePassword{
		ID:          uuid.MustParse(req.Id),
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}

	err := uG.UserService.UpdatePasswordService(input)

	if err != nil {
		return status.Errorf(codes.Internal, "failed to update password : %v", err)
	}

	return nil
}
