package validator

import (
	"context"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/dto"
	"fmt"
)

func (v *validator) CreateEducationalEstablishmentRequest(_ context.Context, dto dto.CreateEducationalEstablishmentRequestDTO) error {
	if len(dto.EducationalEstablishmentName) < 2 {
		return fmt.Errorf("educational establishment name too short: %w", ErrValidation)
	}
	if !v.phoneNumberRegex.MatchString(dto.EducationalEstablishmentContactPhone) {
		return fmt.Errorf("invalid phone number: %w", ErrValidation)
	}
	return nil
}

func (v *validator) PatchEducationalEstablishmentRequest(_ context.Context, dto dto.PatchEducationalEstablishmentRequestDTO) error {
	if len(dto.EducationalEstablishmentName) < 2 {
		return fmt.Errorf("educational establishment name too short: %w", ErrValidation)
	}
	if !v.phoneNumberRegex.MatchString(dto.EducationalEstablishmentContactPhone) {
		return fmt.Errorf("invalid phone number: %w", ErrValidation)
	}
	return nil
}
