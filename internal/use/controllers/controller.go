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
)

type Controller interface {
	Auth(ctx context.Context, token string) (models.User, error)

	SignUp(ctx context.Context, user models.User, request dto.SignUpRequestDTO) (dto.SignUpResponseDTO, error)
	Login(ctx context.Context, request dto.LoginRequestDTO) (dto.LoginResponseDTO, error)
}

type controller struct {
	validator validators.Validator

	repository repositories.Repository
}

func New(validator validators.Validator, repository repositories.Repository) Controller {
	return &controller{validator: validator, repository: repository}
}
