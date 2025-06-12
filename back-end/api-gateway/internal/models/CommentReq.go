package models

type CommentRequest struct {
	UserId  string `json:"userId"`
	PostId  string `json:"postId"`
	Comment string `json:"comment"`
}

type DeleteCommentReq struct {
	Id     string `json:"Id"`
	PostId string `json:"PostId"`
}
