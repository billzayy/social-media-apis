package models

type LoginUserResp struct {
	ID             string `json:"ID"`
	FullName       string `json:"FullName"`
	Email          string `json:"Email"`
	ProfilePicture string `json:"ProfilePicture"`
}

type LoginResp struct {
	User      LoginUserResp `json:"User"`
	Token     string        `json:"Token"`
	Type      string        `json:"Type"`
	ExpiresIn int64         `json:"ExpiresIn"`
}

type SwaggerLoginResp struct {
	StatusCode int       `json:"statusCode"`
	Data       LoginResp `json:"data"`
	Message    string    `json:"message"`
}
