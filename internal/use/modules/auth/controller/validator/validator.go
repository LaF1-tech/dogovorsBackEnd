package validator

import (
	"context"
	"errors"

	"dogovorsBackEnd/internal/use/modules/auth/dto"
)

var (
	ErrValidation = errors.New("validation error")
)

type AuthValidator interface {
	SignUpRequest(ctx context.Context, dto dto.SignUpRequestDTO) error
	LoginRequest(ctx context.Context, dto dto.LoginRequestDTO) error
	UserUpdateRequest(ctx context.Context, dto dto.UserUpdateRequestDTO) error
}

type validator struct {
}

func New() AuthValidator {
	return &validator{}
}
