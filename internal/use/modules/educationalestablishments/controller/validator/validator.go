package validator

import (
	"errors"
	"regexp"
)

var (
	ErrValidation = errors.New("validation error")
)

type EducationalEstablishmentValidator interface {
}

type validator struct {
	phoneNumberRegex *regexp.Regexp
}

func New() EducationalEstablishmentValidator {
	return &validator{
		phoneNumberRegex: regexp.MustCompile(`^\+375(29|44|25|33)[0-9]{7}$`),
	}
}
