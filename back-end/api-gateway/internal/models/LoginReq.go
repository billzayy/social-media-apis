package models

type SwaggerLoginReq struct {
	UserName    string   `json:"UserName"`
	Email       string   `json:"Email"`
	FirstName   string   `json:"FirstName"`
	SurName     string   `json:"SurName"`
	Password    string   `json:"Password"`
	Location    string   `json:"Location"`
	BirthDate   string   `json:"BirthDate"`
	Description string   `json:"Description"`
	Website     []string `json:"Website"`
}
