package models

type LikeRequest struct {
	UserId string `json:"userId"`
	PostId string `json:"postId"`
}

type SwaggerLikeResp struct {
	StatusCode int    `json:"statusCode"`
	Data       bool   `json:"data"`
	Message    string `json:"message"`
}
