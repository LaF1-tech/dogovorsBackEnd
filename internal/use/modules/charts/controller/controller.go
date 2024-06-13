package controller

import (
	"context"

	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/charts/controller/validator"
	"dogovorsBackEnd/internal/use/modules/charts/dto"
	"dogovorsBackEnd/internal/use/modules/charts/repository"
)

type ChartsController interface {
	GetPeriodChart(ctx context.Context, user models.User) (dto.PeriodsChartResponseDTO, error)
	GetPeriodUserChart(ctx context.Context, user models.User, request dto.PeriodUserChartRequestDTO) (dto.PeriodsChartResponseDTO, error)
	GetEducationalEstablishmentChart(ctx context.Context, user models.User) (dto.EducationalEstablishmentsChartResponseDTO, error)
	GetSpecializationsChart(ctx context.Context, user models.User) (dto.SpecializationsChartResponseDTO, error)
	GetTemplatesChart(ctx context.Context, user models.User) (dto.TemplatesChartResponseDTO, error)
}

type controller struct {
	validator  validator.ChartsValidator
	repository repository.ChartsRepository
}

func New(validator validator.ChartsValidator, repository repository.ChartsRepository) ChartsController {
	return &controller{validator: validator, repository: repository}
}
