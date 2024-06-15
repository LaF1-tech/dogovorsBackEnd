package controller

import (
	"context"

	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/contracts/controller/validator"
	"dogovorsBackEnd/internal/use/modules/contracts/dto"
	"dogovorsBackEnd/internal/use/modules/contracts/repository"
)

type ContractsController interface {
	GetAllContracts(ctx context.Context, user models.User) (dto.ContractsViewResponseDTO, error)
	GetContractByID(ctx context.Context, user models.User, id int) (dto.ContractItemResponseDTO, error)

	GenerateContractPDF(ctx context.Context, user models.User, id int) (string, error)
}

type controller struct {
	validator  validator.ContractsValidator
	repository repository.ContractsRepository
}

func New(validator validator.ContractsValidator, repository repository.ContractsRepository) ContractsController {
	return &controller{validator: validator, repository: repository}
}
