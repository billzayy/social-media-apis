package models

type RegisterRequest struct {
	UserName    string   `json:"userName"`
	Email       string   `json:"email"`
	FirstName   string   `json:"firstName"`
	Surname     string   `json:"surName"`
	Password    string   `json:"password"`
	Location    string   `json:"location"`
	BirthDate   string   `json:"birthDate"`
	Description string   `json:"description,omitempty"`
	Website     []string `json:"website,omitempty"`
}
