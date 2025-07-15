package models

type SwaggerPostReq struct {
	UserId  string `json:"UserID"`
	Content string `json:"Content"`
}

type SwaggerPostResp struct {
	StatusCode int        `json:"StatusCode"`
	Data       []PostResp `json:"Data"`
	Message    string     `json:"Message"`
}

type PostResp struct {
	PostId    string      `json:"PostID"`
	Author    Users       `json:"Author"`
	Content   string      `json:"Content"`
	Media     []PostMedia `json:"Media,omitempty"`
	CreatedAt string      `json:"CreatedAt"`
	Likes     int         `json:"Likes"`
	Comments  int         `json:"Comments"`
	Shares    int         `json:"Shares"`
}

type Users struct {
	UserID    string `json:"UserID"`
	FullName  string `json:"FullName"`
	AvatarURL string `json:"ProfilePicture"`
}

type PostMedia struct {
	PostId string `json:"PostID"`
	Type   string `json:"Type"`
	Url    string `json:"Url"`
}

type PostLikes struct {
	UserId   string `json:"UserID"`
	PostId   string `json:"PostID"`
	DateLike string `json:"DateLike"`
	Id       string `json:"ID"`
}

type PostComment struct {
	ID       string `json:"ID"`
	UserId   string `json:"UserId"`
	Comment  string `json:"Comment"`
	SendDate string `json:"SendDate"`
	PostId   string `json:"PostId"`
}
