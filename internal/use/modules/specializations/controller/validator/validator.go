package validator

import (
	"errors"
	"regexp"
)

var (
	ErrValidation = errors.New("validation error")
)

type SpecializationsValidator interface {
}

type validator struct {
	phoneNumberRegex *regexp.Regexp
}

func New() SpecializationsValidator {
	return &validator{
		phoneNumberRegex: regexp.MustCompile(`^\+375(29|44|25|33)[0-9]{7}$`),
	}
}
