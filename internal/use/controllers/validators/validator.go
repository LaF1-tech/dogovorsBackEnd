package validators

import (
	"context"
	"dogovorsBackEnd/internal/use/dto"
	"errors"
)

var (
	ErrValidation = errors.New("validation error")
)

type Validator interface {
	SignUpDTO(ctx context.Context, dto dto.SignUpRequestDTO) error
	LoginDTO(ctx context.Context, dto dto.LoginRequestDTO) error
	UpdateDTO(ctx context.Context, dto dto.UpdateRequestDTO) error
}

type validator struct {
}

func New() Validator {
	return &validator{}
}
