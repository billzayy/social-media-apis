package models

type RefreshUserResp struct {
	ID             string `json:"ID"`
	FullName       string `json:"FullName"`
	Email          string `json:"Email"`
	ProfilePicture string `json:"ProfilePicture"`
}

type RefreshTokenResp struct {
	User      RefreshUserResp `json:"User"`
	Token     string          `json:"Token"`
	Type      string          `json:"Type"`
	ExpiresIn int64           `json:"ExpiresIn"`
}
