package models

type LikeRequest struct {
	UserId string `json:"UserID"`
	PostId string `json:"PostID"`
}

type SwaggerLikeResp struct {
	StatusCode int    `json:"StatusCode"`
	Data       bool   `json:"Data"`
	Message    string `json:"Message"`
}

type CommentRequest struct {
	UserId  string `json:"UserID"`
	PostId  string `json:"PostID"`
	Comment string `json:"Comment"`
}

type DeleteCommentReq struct {
	Id     string `json:"ID"`
	PostId string `json:"PostId"`
}
