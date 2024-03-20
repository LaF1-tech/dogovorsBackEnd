package validator

import (
	"errors"
)

var (
	ErrValidation = errors.New("validation error")
)

type ContractsValidator interface {
}

type validator struct {
}

func New() ContractsValidator {
	return &validator{}
}
