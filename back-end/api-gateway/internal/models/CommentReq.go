package models

type CommentRequest struct {
	UserId  string `json:"userId"`
	PostId  string `json:"postId"`
	Comment string `json:"comment"`
}
