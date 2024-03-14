package validator

import (
	"context"
	"errors"
	"regexp"

	"dogovorsBackEnd/internal/use/modules/applications/dto"
)

var (
	ErrValidation = errors.New("validation error")
)

type ApplicationValidator interface {
	CreateApplicationRequest(ctx context.Context, request dto.CreateApplicationRequestDTO) error
}

type validator struct {
	phoneNumberRegex *regexp.Regexp
}

func New() ApplicationValidator {
	return &validator{
		phoneNumberRegex: regexp.MustCompile(`^\+375(29|44|25|33)[0-9]{7}$`),
	}
}
