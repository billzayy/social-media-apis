package models

type LikeRequest struct {
	UserId string `json:"userId"`
	PostId string `json:"postId"`
}
