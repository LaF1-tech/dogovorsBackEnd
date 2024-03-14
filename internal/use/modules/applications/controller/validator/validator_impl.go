package validator

import (
	"context"
	"fmt"

	"dogovorsBackEnd/internal/use/modules/applications/dto"
)

func (v *validator) CreateApplicationRequest(_ context.Context, dto dto.CreateApplicationRequestDTO) error {
	if len(dto.Name) < 2 {
		return fmt.Errorf("name is too short: %w", ErrValidation)
	}
	if len(dto.LastName) < 2 {
		return fmt.Errorf("last name is too short: %w", ErrValidation)
	}
	if len(dto.MiddleName) < 2 {
		return fmt.Errorf("middle name is too short: %w", ErrValidation)
	}
	if !v.phoneNumberRegex.MatchString(dto.PhoneNumber) {
		return fmt.Errorf("invalid phone number: %w", ErrValidation)
	}

	return nil
}
