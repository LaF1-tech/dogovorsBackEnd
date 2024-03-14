package controller

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/specializations/controller/validator"
	"dogovorsBackEnd/internal/use/modules/specializations/dto"
	"dogovorsBackEnd/internal/use/modules/specializations/repository"
)

type SpecializationsController interface {
	GetSpecializations(ctx context.Context) (dto.SpecializationsResponseDTO, error)
}

type controller struct {
	validator  validator.SpecializationsValidator
	repository repository.SpecializationsRepository
}

func New(validator validator.SpecializationsValidator, repository repository.SpecializationsRepository) SpecializationsController {
	return &controller{validator: validator, repository: repository}
}
