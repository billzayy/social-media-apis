package handlers

import (
	"context"
	"log"

	grpc "github.com/billzayy/social-media/back-end/post-service/api"
	"github.com/billzayy/social-media/back-end/post-service/internal/models"
	"github.com/billzayy/social-media/back-end/post-service/internal/services"
	"github.com/billzayy/social-media/back-end/post-service/internal/utils"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	grpc.UnimplementedPostServiceServer
	PostService     *services.PostService
	InteractService *services.InteractService
}

func NewPostGrpcServer(sv *services.Services) *GrpcServer {
	return &GrpcServer{
		PostService:     sv.PostService,
		InteractService: sv.InteractService,
	}
}

func (pG *GrpcServer) GetPost(ctx context.Context, _ *emptypb.Empty) (*grpc.GetPostResp, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in GetPost")
		}
	}()

	posts, err := pG.PostService.GetPost()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	var protoPosts []*grpc.PostResp
	for _, p := range posts {
		protoPosts = append(protoPosts, &grpc.PostResp{
			PostId:    p.PostId.String(),
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			Likes:     int32(p.Likes),
			Comments:  int32(p.Comments),
			Shares:    int32(p.Shares),
			Author: &grpc.Users{
				UserId:         p.Author.UserId.String(),
				FullName:       p.Author.FullName,
				ProfilePicture: p.Author.AvatarURL,
			},
			// Convert media if needed:
			Media: utils.ConvertMedia(p.Media),
		})
	}

	return &grpc.GetPostResp{PostList: protoPosts}, nil
}

func (pG *GrpcServer) CreatePost(ctx context.Context, request *grpc.CreatePostReq) (*grpc.CreatePostResp, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in CreatePost")
		}
	}()

	r := models.AddPostRequest{
		UserId:    uuid.MustParse(request.UserId),
		Content:   request.Content,
		CreatedAt: request.CreatedAt,
	}

	valid, err := pG.PostService.CreatePost(r)

	if err != nil {
		return &grpc.CreatePostResp{Message: "Create Failed !"}, status.Errorf(codes.Internal, "%v", err)
	}

	if valid == false {
		return &grpc.CreatePostResp{Message: "Create Failed"}, status.Errorf(codes.Internal, "%v", err)
	}

	return &grpc.CreatePostResp{Message: "Post Created !"}, nil
}

func (pG *GrpcServer) DeletePost(ctx context.Context, req *grpc.DeletePostReq) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in DeletePost")
		}
	}()

	err := pG.PostService.DeletePost(req.Id)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	return &emptypb.Empty{}, nil
}

func (iG *GrpcServer) CheckLike(ctx context.Context, req *grpc.LikeRequest) (*grpc.CheckLikeResp, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in CheckLike")
		}
	}()

	userId, err := uuid.Parse(req.UserId)

	if err != nil {
		return &grpc.CheckLikeResp{Valid: false}, status.Errorf(codes.Internal, "%v", err)
	}

	postId, err := uuid.Parse(req.PostId)

	if err != nil {
		return &grpc.CheckLikeResp{Valid: false}, status.Errorf(codes.Internal, "%v", err)
	}

	valid, err := iG.InteractService.CheckLikeOnPostService(userId, postId)

	if err != nil {
		return &grpc.CheckLikeResp{Valid: false}, status.Errorf(codes.Internal, "%v", err)
	}

	if valid == false {
		return &grpc.CheckLikeResp{Valid: false}, nil
	}

	return &grpc.CheckLikeResp{Valid: valid}, nil
}

func (iG *GrpcServer) AddLike(ctx context.Context, req *grpc.LikeRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in AddLike")
		}
	}()

	userId, err := uuid.Parse(req.UserId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	postId, err := uuid.Parse(req.PostId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	err = iG.InteractService.AddLikeService(userId, postId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	return &emptypb.Empty{}, nil
}

func (iG *GrpcServer) RemoveLike(ctx context.Context, req *grpc.LikeRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in RemoveLike")
		}
	}()

	userId, err := uuid.Parse(req.UserId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	postId, err := uuid.Parse(req.PostId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	err = iG.InteractService.RemoveLikeService(userId, postId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	return &emptypb.Empty{}, nil
}

func (iG *GrpcServer) AddComment(ctx context.Context, req *grpc.AddCommentReq) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in Add Comment")
		}
	}()

	userId, err := uuid.Parse(req.UserId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	postId, err := uuid.Parse(req.PostId)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	input := models.CommentRequest{
		UserId:  userId,
		PostId:  postId,
		Comment: req.Comment,
	}

	err = iG.InteractService.AddCommentService(input)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	return &emptypb.Empty{}, nil
}

func (iG *GrpcServer) DeleteComment(ctx context.Context, req *grpc.DeleteCommentReq) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in DeleteComment")
		}
	}()

	id, err := uuid.Parse(req.Id)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	err = iG.InteractService.DeleteCommentService(id)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "%v", err)
	}

	return &emptypb.Empty{}, nil
}
