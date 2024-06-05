package validator

import (
	"context"
	"dogovorsBackEnd/internal/use/modules/specializations/dto"
	"fmt"
)

func (v *validator) CreateSpecializationRequest(ctx context.Context, request dto.CreateSpecializationRequestDTO) error {
	if len(request.SpecializationName) < 2 {
		return fmt.Errorf("specialization name is too short, %w", ErrValidation)
	}
	return nil
}

func (v *validator) PatchSpecializationRequest(ctx context.Context, request dto.PatchSpecializationRequestDTO) error {
	if len(request.SpecializationName) < 2 {
		return fmt.Errorf("specialization name is too short, %w", ErrValidation)
	}
	return nil
}
