package validator

import (
	"context"
	"fmt"

	"dogovorsBackEnd/internal/use/modules/auth/dto"
)

func (v *validator) SignUpRequest(_ context.Context, dto dto.SignUpRequestDTO) error {
	if dto.FirstName == "" {
		return fmt.Errorf("first name is empty: %w", ErrValidation)
	}
	if dto.LastName == "" {
		return fmt.Errorf("last name is empty: %w", ErrValidation)
	}

	if len(dto.Username) < 3 {
		return fmt.Errorf("username is too short, min length is 3: %w", ErrValidation)
	}
	if len(dto.Password) < 4 {
		return fmt.Errorf("password is too short, min length is 4: %w", ErrValidation)
	}

	return nil
}

func (v *validator) LoginRequest(_ context.Context, dto dto.LoginRequestDTO) error {
	if len(dto.Username) < 3 {
		return fmt.Errorf("username is too short, min length is 3: %w", ErrValidation)
	}
	if len(dto.Password) < 4 {
		return fmt.Errorf("password is too short, min length is 4: %w", ErrValidation)
	}

	return nil
}

func (v *validator) UserUpdateRequest(_ context.Context, dto dto.UserUpdateRequestDTO) error {
	if len(dto.Password) < 4 && len(dto.Password) != 0 {
		return fmt.Errorf("password is too short, min length is 4: %w", ErrValidation)
	}

	return nil
}
