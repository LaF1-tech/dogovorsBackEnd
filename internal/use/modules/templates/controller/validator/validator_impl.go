package validator

import (
	"context"
	"dogovorsBackEnd/internal/use/modules/templates/dto"
	"fmt"
)

func (v *validator) CreateTemplateRequest(_ context.Context, dto dto.CreateTemplateRequestDTO) error {
	if len(dto.TemplateName) < 2 {
		return fmt.Errorf("template name is too short: %w", ErrValidation)
	}
	if len(dto.TemplateContent) < 10 {
		return fmt.Errorf("template content is too short: %w", ErrValidation)
	}
	return nil
}

func (v *validator) PatchTemplateRequest(_ context.Context, dto dto.PatchTemplateRequestDTO) error {
	if len(dto.TemplateName) < 2 {
		return fmt.Errorf("template name is too short: %w", ErrValidation)
	}
	if len(dto.TemplateContent) < 10 {
		return fmt.Errorf("template content is too short: %w", ErrValidation)
	}
	return nil
}
