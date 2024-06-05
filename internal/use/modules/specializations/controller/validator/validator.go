package validator

import (
	"context"
	"dogovorsBackEnd/internal/use/modules/specializations/dto"
	"errors"
	"regexp"
)

var (
	ErrValidation = errors.New("validation error")
)

type SpecializationsValidator interface {
	CreateSpecializationRequest(ctx context.Context, request dto.CreateSpecializationRequestDTO) error
	PatchSpecializationRequest(ctx context.Context, request dto.PatchSpecializationRequestDTO) error
}

type validator struct {
	phoneNumberRegex *regexp.Regexp
}

func New() SpecializationsValidator {
	return &validator{
		phoneNumberRegex: regexp.MustCompile(`^\+375(29|44|25|33)[0-9]{7}$`),
	}
}
