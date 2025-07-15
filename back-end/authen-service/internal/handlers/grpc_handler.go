package handlers

import (
	"context"
	"log"

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
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in Register")
		}
	}()

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
		return &auth.RegisterResp{Successful: false}, status.Errorf(codes.Internal, "%v", err)
	}

	return &auth.RegisterResp{Successful: data}, nil
}

func (aG *AuthGrpcServer) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in Login")
		}
	}()

	data, cookie, err := aG.AuthService.LoginService(req.GetUserName(), req.GetPassword())
	if err != nil {
		return &auth.LoginResp{}, status.Errorf(codes.Internal, "%v", err)
	}

	return &auth.LoginResp{
		Token:   data.Token,
		Type:    data.Type,
		Cookie:  cookie.Value,
		Expires: data.ExpiresIn,
		User: &auth.UserResp{
			ID:             data.User.ID,
			FullName:       data.User.FullName,
			Email:          data.User.Email,
			ProfilePicture: data.User.ProfilePicture,
		},
	}, nil
}

func (aG *AuthGrpcServer) RefreshToken(ctx context.Context, req *auth.RefreshTokenReq) (*auth.RefreshTokenResp, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in RefreshToken")
		}
	}()

	userData, newToken, expires, cookie, err := aG.AuthService.RefreshTokenService(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &auth.RefreshTokenResp{
		Token:   newToken,
		Type:    "Bearer",
		Cookie:  cookie.Value,
		Expires: expires,
		User: &auth.UserResp{
			ID:             userData.ID,
			FullName:       userData.FullName,
			Email:          userData.Email,
			ProfilePicture: userData.ProfilePicture,
		},
	}, nil
}
