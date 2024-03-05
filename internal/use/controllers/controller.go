package controllers

import (
	"context"
	"dogovorsBackEnd/internal/use/controllers/validators"
	"dogovorsBackEnd/internal/use/dto"
	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/repositories"
	"errors"
)

var (
	ErrInvalidPassword = errors.New("invalid password error")
	ErrForbidden       = errors.New("forbidden")
)

type Controller interface {
	Auth(ctx context.Context, token string) (models.User, error)

	SignUp(ctx context.Context, user models.User, request dto.SignUpRequestDTO) (dto.SignUpResponseDTO, error)
	Login(ctx context.Context, request dto.LoginRequestDTO) (dto.LoginResponseDTO, error)
	PatchUser(ctx context.Context, user models.User, request dto.UpdateRequestDTO) (dto.UpdateResponseDTO, error)
	DeleteUserByID(ctx context.Context, user models.User, userID int) error
}

type controller struct {
	validator validators.Validator

	repository repositories.Repository
}

func New(validator validators.Validator, repository repositories.Repository) Controller {
	return &controller{validator: validator, repository: repository}
}
