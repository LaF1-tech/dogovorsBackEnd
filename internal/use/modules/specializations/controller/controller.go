package controller

import (
	"context"
	"dogovorsBackEnd/internal/use/models"

	"dogovorsBackEnd/internal/use/modules/specializations/controller/validator"
	"dogovorsBackEnd/internal/use/modules/specializations/dto"
	"dogovorsBackEnd/internal/use/modules/specializations/repository"
)

type SpecializationsController interface {
	CreateSpecialization(ctx context.Context, user models.User, request dto.CreateSpecializationRequestDTO) (dto.SpecializationItemResponseDTO, error)
	GetSpecializations(ctx context.Context) (dto.SpecializationsResponseDTO, error)
	GetSpecializationByID(ctx context.Context, id int) (dto.SpecializationItemResponseDTO, error)
	PatchSpecializationByID(ctx context.Context, user models.User, request dto.PatchSpecializationRequestDTO) error
	DeleteSpecializationByID(ctx context.Context, user models.User, id int) error
}

type controller struct {
	validator  validator.SpecializationsValidator
	repository repository.SpecializationsRepository
}

func New(validator validator.SpecializationsValidator, repository repository.SpecializationsRepository) SpecializationsController {
	return &controller{validator: validator, repository: repository}
}
