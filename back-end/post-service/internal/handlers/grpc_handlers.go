package handlers

import (
	"context"

	post "github.com/billzayy/social-media/back-end/post-service/api"
	"github.com/billzayy/social-media/back-end/post-service/internal/models"
	"github.com/billzayy/social-media/back-end/post-service/internal/services"
	"github.com/billzayy/social-media/back-end/post-service/internal/utils"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostGrpcServer struct {
	post.UnimplementedPostServiceServer
	PostService     *services.PostService
	InteractService *services.InteractService
}

func NewPostGrpcServer(sv *services.Services) *PostGrpcServer {
	return &PostGrpcServer{
		PostService:     sv.PostService,
		InteractService: sv.InteractService,
	}
}

func (pG *PostGrpcServer) GetPost(ctx context.Context, _ *emptypb.Empty) (*post.GetPostResp, error) {
	posts, err := pG.PostService.GetPost()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to Get post : %v", err)
	}

	var protoPosts []*post.PostResp
	for _, p := range posts {
		protoPosts = append(protoPosts, &post.PostResp{
			PostId:    p.PostId.String(),
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			Likes:     int32(p.Likes),
			Comments:  int32(p.Comments),
			Shares:    int32(p.Shares),
			Author: &post.Users{
				UserId:         p.Author.UserId.String(),
				FullName:       p.Author.FullName,
				ProfilePicture: p.Author.AvatarURL,
			},
			// Convert media if needed:
			Media: utils.ConvertMedia(p.Media),
		})
	}

	return &post.GetPostResp{PostList: protoPosts}, nil
}

func (pG *PostGrpcServer) CreatePost(ctx context.Context, request *post.CreatePostReq) (*post.CreatePostResp, error) {

	r := models.AddPostRequest{
		UserId:    uuid.MustParse(request.UserId),
		Content:   request.Content,
		CreatedAt: request.CreatedAt,
	}

	valid, err := pG.PostService.CreatePost(r)

	if err != nil {
		return &post.CreatePostResp{Message: "Create Failed !"}, err
	}

	if valid == false {
		return &post.CreatePostResp{Message: "Create Failed"}, err
	}

	return &post.CreatePostResp{Message: "Post Created !"}, nil
}

func (pG *PostGrpcServer) DeletePost(ctx context.Context, req *post.DeletePostReq) (*emptypb.Empty, error) {
	err := pG.PostService.DeletePost(req.Id)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
