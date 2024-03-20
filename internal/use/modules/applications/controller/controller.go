package controller

import (
	"context"

	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/applications/controller/validator"
	"dogovorsBackEnd/internal/use/modules/applications/dto"
	"dogovorsBackEnd/internal/use/modules/applications/repository"
)

type ApplicationsController interface {
	CreateApplication(ctx context.Context, request dto.CreateApplicationRequestDTO) (dto.ApplicationItemResponseDTO, error)
	GetAllApplications(ctx context.Context, user models.User) (dto.ApplicationsViewResponseDTO, error)
	PatchApplicationByID(ctx context.Context, user models.User, request dto.PatchApplicationRequestDTO) error
}

type controller struct {
	validator  validator.ApplicationValidator
	repository repository.ApplicationsRepository
}

func New(validator validator.ApplicationValidator, repository repository.ApplicationsRepository) ApplicationsController {
	return &controller{validator: validator, repository: repository}
}
