package controller

import (
	"context"
	"errors"

	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/auth/controller/validator"
	"dogovorsBackEnd/internal/use/modules/auth/dto"
	"dogovorsBackEnd/internal/use/modules/auth/repository"
)

var (
	ErrInvalidPassword = errors.New("invalid password error")
	ErrForbidden       = errors.New("forbidden")
)

type AuthController interface {
	Auth(ctx context.Context, token string) (models.User, error)

	SignUp(ctx context.Context, user models.User, request dto.SignUpRequestDTO) (dto.SignUpResponseDTO, error)
	Login(ctx context.Context, request dto.LoginRequestDTO) (dto.LoginResponseDTO, error)
	PatchUser(ctx context.Context, user models.User, request dto.UserUpdateRequestDTO) (dto.UserUpdateResponseDTO, error)
	DeleteUserByID(ctx context.Context, user models.User, userID int) error
}

type controller struct {
	validator validator.AuthValidator

	repository repository.AuthRepository
}

func New(validator validator.AuthValidator, repository repository.AuthRepository) AuthController {
	return &controller{validator: validator, repository: repository}
}
