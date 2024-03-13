package dto

import "dogovorsBackEnd/internal/use/models"

type SignUpRequestDTO struct {
	Username    string              `json:"username,omitempty"`
	Password    string              `json:"password,omitempty"`
	FirstName   string              `json:"first_name,omitempty"`
	LastName    string              `json:"last_name,omitempty"`
	Permissions []models.Permission `json:"permissions,omitempty"`
}

type LoginRequestDTO struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserUpdateRequestDTO struct {
	UserID      int                 `json:"user_id,omitempty"`
	Password    string              `json:"password,omitempty"`
	FirstName   string              `json:"first_name,omitempty"`
	LastName    string              `json:"last_name,omitempty"`
	Permissions []models.Permission `json:"permissions,omitempty"`
}
