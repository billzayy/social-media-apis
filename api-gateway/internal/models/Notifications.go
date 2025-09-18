package models

type SwaggerReqSendNotify struct {
	SenderId   string `json:"SenderId"`
	ReceiverId string `json:"ReceiverId"`
	Messages   string `json:"Messages"`
	Type       string `json:"Type"`
	Url        string `json:"Url"`
}

type SwaggerRespGetList struct {
	Notifications []*Notifications `json:"Notifications"`
}

type Notifications struct {
	Id         string `json:"Id"`
	SenderId   string `json:"SenderId"`
	ReceiverId string `json:"ReceiverId"`
	Messages   string `json:"Messages"`
	Type       string `json:"Type"`
	Url        string `json:"Url"`
	IsRead     bool   `json:"IsRead"`
	Date       int64  `json:"Date"`
}
