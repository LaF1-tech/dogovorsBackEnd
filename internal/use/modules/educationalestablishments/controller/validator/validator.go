package validator

import (
	"context"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/dto"
	"errors"
	"regexp"
)

var (
	ErrValidation = errors.New("validation error")
)

type EducationalEstablishmentValidator interface {
	CreateEducationalEstablishmentRequest(ctx context.Context, dto dto.CreateEducationalEstablishmentRequestDTO) error
	PatchEducationalEstablishmentRequest(ctx context.Context, dto dto.PatchEducationalEstablishmentRequestDTO) error
}

type validator struct {
	phoneNumberRegex *regexp.Regexp
}

func New() EducationalEstablishmentValidator {
	return &validator{
		phoneNumberRegex: regexp.MustCompile(`^\+375(29|44|25|33)[0-9]{7}$`),
	}
}
