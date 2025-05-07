package handlers

import (
	"context"
	"fmt"

	post "github.com/billzayy/social-media/back-end/post-service/api"
	"github.com/billzayy/social-media/back-end/post-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/post-service/internal/models"
	"github.com/billzayy/social-media/back-end/post-service/internal/utils"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostGrpcServer struct {
	post.UnimplementedPostServiceServer
	PostRepository     *repositories.PostRepository
	InteractRepository *repositories.InteractRepository
}

func NewPostGrpcServer(dbPost *repositories.PostRepository, dbInteract *repositories.InteractRepository) *PostGrpcServer {
	return &PostGrpcServer{
		PostRepository:     dbPost,
		InteractRepository: dbInteract,
	}
}

func (pG *PostGrpcServer) GetPost(ctx context.Context, _ *emptypb.Empty) (*post.GetPostResp, error) {
	posts, err := pG.PostRepository.GetPost()

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

func (pG *PostGrpcServer) AddPost(ctx context.Context, request *post.AddPostReq) (*post.AddPostResp, error) {

	r := models.AddPostRequest{
		UserId:    uuid.MustParse(request.UserId),
		Content:   request.Content,
		CreatedAt: request.CreatedAt,
	}

	err := pG.PostRepository.AddPost(r)

	if err != nil {
		return &post.AddPostResp{Message: "Create Failed !"}, err
	}

	return &post.AddPostResp{Message: "Post Created !"}, nil
}

func (pG *PostGrpcServer) DeletePost(ctx context.Context, req *post.DeletePostReq) (*emptypb.Empty, error) {
	resp, err := pG.PostRepository.DeletePost(req.Id)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	if resp == 0 {
		return &emptypb.Empty{}, fmt.Errorf("post %s not found", req.Id)
	}

	return &emptypb.Empty{}, nil
}
