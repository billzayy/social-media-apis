package handlers

import (
	"context"

	auth "github.com/billzayy/social-media/back-end/authen-service/api"
	"github.com/billzayy/social-media/back-end/authen-service/internal/models"
	"github.com/billzayy/social-media/back-end/authen-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthGrpcServer struct {
	auth.UnimplementedAuthServiceServer
	AuthService *services.AuthService
}

func NewAuthGrpcServer(sv *services.AuthService) *AuthGrpcServer {
	return &AuthGrpcServer{AuthService: sv}
}

func (aG *AuthGrpcServer) Register(ctx context.Context, req *auth.RegisterReq) (*auth.RegisterResp, error) {
	requestData := models.RegisterRequest{
		UserName:    req.GetUserName(),
		Email:       req.GetEmail(),
		FirstName:   req.GetFirstName(),
		Surname:     req.GetSurName(),
		Password:    req.GetPassword(),
		Location:    req.GetLocation(),
		BirthDate:   req.GetBirthDate(),
		Description: req.GetDescription(),
		Website:     req.GetWebsite(),
	}

	data, err := aG.AuthService.RegisterService(requestData)

	if err != nil || !data {
		return &auth.RegisterResp{Successful: false, ErrorMessage: err.Error()}, status.Errorf(codes.Internal, "failed to register: %v", err)
	}

	return &auth.RegisterResp{Successful: data}, nil
}

func (aG *AuthGrpcServer) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	token, cookie, err := aG.AuthService.LoginService(req.GetUserName(), req.GetPassword())
	if err != nil {
		return &auth.LoginResp{ErrorMessage: err.Error()}, status.Errorf(codes.Internal, "login failed: %v", err)
	}

	return &auth.LoginResp{
		Token:  token.Token,
		Type:   token.Type,
		Cookie: cookie.Value,
		UserId: token.UserId.String(),
	}, nil
}

func (aG *AuthGrpcServer) RefreshToken(ctx context.Context, req *auth.RefreshTokenReq) (*auth.RefreshTokenResp, error) {
	userId, newToken, cookie, err := aG.AuthService.RefreshTokenService(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "refresh token failed: %v", err)
	}

	return &auth.RefreshTokenResp{
		Token:  newToken,
		Type:   "Bearer",
		Cookie: cookie.Value,
		UserId: userId,
	}, nil
}
