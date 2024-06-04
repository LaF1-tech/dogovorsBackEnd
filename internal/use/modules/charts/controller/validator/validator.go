package validator

import (
	"errors"
)

var (
	ErrValidation = errors.New("validation error")
)

type ChartsValidator interface {
}

type validator struct {
}

func New() ChartsValidator {
	return &validator{}
}
