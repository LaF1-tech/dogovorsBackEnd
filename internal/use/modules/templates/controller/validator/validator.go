package validator

import (
	"context"
	"dogovorsBackEnd/internal/use/modules/templates/dto"
	"errors"
)

var (
	ErrValidation = errors.New("validation error")
)

type TemplateValidator interface {
	CreateTemplateRequest(_ context.Context, dto dto.CreateTemplateRequestDTO) error
	PatchTemplateRequest(_ context.Context, dto dto.PatchTemplateRequestDTO) error
}

type validator struct {
}

func New() TemplateValidator {
	return &validator{}
}
