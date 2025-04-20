package handlers

import (
	"context"

	auth "github.com/billzayy/social-media/back-end/authen-service/api"
	"github.com/billzayy/social-media/back-end/authen-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/authen-service/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthGrpcServer struct {
	auth.UnimplementedAuthServiceServer
	AuthRepository *repositories.AuthRepository
}

func NewAuthGrpcServer(db *repositories.AuthRepository) *AuthGrpcServer {
	return &AuthGrpcServer{AuthRepository: db}
}

func (aG *AuthGrpcServer) Register(ctx context.Context, req *auth.RegisterReq) (*auth.RegisterResp, error) {
	requestData := models.RegisterRequest{
		UserName:    req.UserName,
		Email:       req.Email,
		FirstName:   req.FirstName,
		Surname:     req.SurName,
		Password:    req.Password,
		Location:    req.Location,
		BirthDate:   req.BirthDate,
		Description: req.Description,
		Website:     req.Website,
	}

	err := aG.AuthRepository.Register(requestData)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register: %v", err)
	}

	return &auth.RegisterResp{Message: "Account Created!"}, nil
}

func (aG *AuthGrpcServer) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	hashed, checkData, err := aG.AuthRepository.CheckUserWithNameAndPass(req.UserName, req.Password)

	if err != nil || !checkData {
		return nil, status.Errorf(codes.InvalidArgument, "invalid credentials")
	}

	token, _, err := aG.AuthRepository.Login(req.UserName, hashed)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "login failed: %v", err)
	}

	return &auth.LoginResp{
		Token:  token.Token,
		Type:   token.Type,
		UserId: token.UserId.String(),
	}, nil
}

func (aG *AuthGrpcServer) RefreshToken(ctx context.Context, req *auth.RefreshTokenReq) (*auth.RefreshTokenResp, error) {
	userId, newToken, _, err := aG.AuthRepository.RefreshToken(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "refresh token failed: %v", err)
	}

	return &auth.RefreshTokenResp{
		Token:  newToken,
		Type:   "Bearer",
		UserId: userId,
	}, nil
}
