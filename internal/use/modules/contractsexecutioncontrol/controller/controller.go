package controller

import (
	"context"
	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/controller/validator"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/dto"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/repository"
)

type ContractsExecutionControlController interface {
	GetContractsExecutionControlByContractID(ctx context.Context, user models.User, id int) (dto.AggregatedContractsExecutionControlResponseDTO, error)
	CreateContractExecutionControl(ctx context.Context, user models.User, request dto.CreateContractExecutionControlRequestDTO) (dto.ContractExecutionControlResponseDTO, error)
	PatchContractExecutionControl(ctx context.Context, user models.User, request dto.PatchContractExecutionControlRequestDTO) error
}

type controller struct {
	validator  validator.ContractsExecutionControlValidator
	repository repository.ContractsExecutionControlRepository
}

func New(validator validator.ContractsExecutionControlValidator, repository repository.ContractsExecutionControlRepository) ContractsExecutionControlController {
	return &controller{validator: validator, repository: repository}
}
