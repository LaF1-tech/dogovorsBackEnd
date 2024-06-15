package validator

import (
	"context"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/dto"
	"fmt"
)

func (v *validator) CreateContractExecutionControlRequest(_ context.Context, dto dto.CreateContractExecutionControlRequestDTO) error {
	if len(dto.ControlStatus) < 5 {
		return fmt.Errorf("control status is too short: %w", ErrValidation)
	}
	return nil
}

func (v *validator) PatchContractExecutionControlRequest(_ context.Context, dto dto.PatchContractExecutionControlRequestDTO) error {
	if len(dto.ControlStatus) < 5 {
		return fmt.Errorf("control status is too short: %w", ErrValidation)
	}
	return nil
}
