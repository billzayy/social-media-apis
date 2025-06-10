package models

type SwaggerPostReq struct {
	UserId  string `json:"userId"`
	Content string `json:"content"`
}

type SwaggerPostResp struct {
	StatusCode int        `json:"statusCode"`
	Data       []PostResp `json:"data"`
	Message    string     `json:"message"`
}

type PostResp struct {
	PostId    string      `json:"postId"`
	Author    Users       `json:"userId"`
	Content   string      `json:"content"`
	Media     []PostMedia `json:"media,omitempty"`
	CreatedAt string      `json:"createdAt"`
	Likes     int         `json:"likes"`
	Comments  int         `json:"comments"`
	Shares    int         `json:"shares"`
}

type Users struct {
	UserId    string `json:"ID"`
	FullName  string `json:"fullName"`
	AvatarURL string `json:"profilePicture"`
}

type PostMedia struct {
	PostId string `json:"postId"`
	Type   string `json:"type"`
	Url    string `json:"url"`
}

type PostLikes struct {
	UserId   string `json:"UserID"`
	PostId   string `json:"PostId"`
	DateLike string `json:"DateLike"`
	Id       string `json:"id"`
}

type PostComment struct {
	ID       string `json:"ID"`
	UserId   string `json:"UserId"`
	Comment  string `json:"Comment"`
	SendDate string `json:"SendDate"`
	PostId   string `json:"PostId"`
}
