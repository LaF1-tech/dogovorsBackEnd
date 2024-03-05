package validators

import (
	"context"
	"dogovorsBackEnd/internal/use/dto"
	"fmt"
)

func (v *validator) SignUpDTO(_ context.Context, dto dto.SignUpRequestDTO) error {
	if dto.FirstName == "" {
		return fmt.Errorf("first name is empty: %w", ErrValidation)
	}
	if dto.LastName == "" {
		return fmt.Errorf("last name is empty: %w", ErrValidation)
	}

	if len(dto.Username) < 3 {
		return fmt.Errorf("username is too short, min length is 6: %w", ErrValidation)
	}
	if len(dto.Password) < 4 {
		return fmt.Errorf("password is too short, min length is 4: %w", ErrValidation)
	}

	return nil
}

func (v *validator) LoginDTO(_ context.Context, dto dto.LoginRequestDTO) error {
	if len(dto.Username) < 3 {
		return fmt.Errorf("username is too short, min length is 6: %w", ErrValidation)
	}
	if len(dto.Password) < 4 {
		return fmt.Errorf("password is too short, min length is 4: %w", ErrValidation)
	}

	return nil
}

func (v *validator) UpdateDTO(_ context.Context, dto dto.UpdateRequestDTO) error {
	if len(dto.Password) < 4 && len(dto.Password) != 0 {
		return fmt.Errorf("password is too short, min length is 4: %w", ErrValidation)
	}

	return nil
}
