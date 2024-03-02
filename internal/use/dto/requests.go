package dto

type SignUpRequestDTO struct {
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type LoginRequestDTO struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
