package validator

import (
	"context"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/dto"
	"errors"
)

var (
	ErrValidation = errors.New("validation error")
)

type ContractsExecutionControlValidator interface {
	CreateContractExecutionControlRequest(_ context.Context, dto dto.CreateContractExecutionControlRequestDTO) error
	PatchContractExecutionControlRequest(_ context.Context, dto dto.PatchContractExecutionControlRequestDTO) error
}

type validator struct {
}

func New() ContractsExecutionControlValidator {
	return &validator{}
}
