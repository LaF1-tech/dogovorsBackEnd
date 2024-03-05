package dto

import "dogovorsBackEnd/internal/use/models"

type SignUpResponseDTO struct {
	UserID      int                 `json:"user_id,omitempty"`
	Username    string              `json:"username,omitempty"`
	FirstName   string              `json:"first_name,omitempty"`
	LastName    string              `json:"last_name,omitempty"`
	Permissions []models.Permission `json:"permissions,omitempty"`
}

type LoginUserResponseDTO struct {
	UserID      int                 `json:"user_id,omitempty"`
	Username    string              `json:"username,omitempty"`
	FirstName   string              `json:"first_name,omitempty"`
	LastName    string              `json:"last_name,omitempty"`
	Permissions []models.Permission `json:"permissions,omitempty"`
}

type LoginResponseDTO struct {
	User  LoginUserResponseDTO `json:"user"`
	Token string               `json:"token"`
}

type UpdateResponseDTO struct {
	UserID      int                 `json:"user_id,omitempty"`
	Username    string              `json:"username,omitempty"`
	FirstName   string              `json:"first_name,omitempty"`
	LastName    string              `json:"last_name,omitempty"`
	Permissions []models.Permission `json:"permissions,omitempty"`
}
