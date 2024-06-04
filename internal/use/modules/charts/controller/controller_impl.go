package controller

import (
	"context"
	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/charts/dto"
	"dogovorsBackEnd/internal/use/modules/charts/entities"
	"dogovorsBackEnd/internal/use/utils/uslices"
)

func (c *controller) GetPeriodChart(ctx context.Context, user models.User) (dto.PeriodsChartResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.PeriodsChartResponseDTO{}, err
	}
	contractsChart, err := c.repository.GetPeriodChart(ctx)
	if err != nil {
		return dto.PeriodsChartResponseDTO{}, err
	}
	return dto.PeriodsChartResponseDTO{
		List: uslices.MapFunc(contractsChart, func(item entities.PeriodChart) dto.PeriodChartResponseDTO {
			return dto.PeriodChartResponseDTO{
				Period:        item.Period,
				ContractCount: item.ContractCount,
			}
		}),
	}, nil
}

func (c *controller) GetEducationalEstablishmentChart(ctx context.Context, user models.User) (dto.EducationalEstablishmentsChartResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.EducationalEstablishmentsChartResponseDTO{}, err
	}
	educationalEstablishmentChart, err := c.repository.GetEducationalEstablishmentChart(ctx)
	if err != nil {
		return dto.EducationalEstablishmentsChartResponseDTO{}, err
	}
	return dto.EducationalEstablishmentsChartResponseDTO{
		List: uslices.MapFunc(educationalEstablishmentChart, func(item entities.EducationalEstablishmentCountChart) dto.EducationalEstablishmentChartResponseDTO {
			return dto.EducationalEstablishmentChartResponseDTO{
				EducationalEstablishmentName: item.EducationalEstablishmentName,
				ContractCount:                item.ContractCount,
			}
		}),
	}, nil
}

func (c *controller) GetSpecializationsChart(ctx context.Context, user models.User) (dto.SpecializationsChartResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.SpecializationsChartResponseDTO{}, err
	}
	specializationsChart, err := c.repository.GetSpecializationsChart(ctx)
	if err != nil {
		return dto.SpecializationsChartResponseDTO{}, err
	}
	return dto.SpecializationsChartResponseDTO{
		List: uslices.MapFunc(specializationsChart, func(item entities.SpecializationNameCountChart) dto.SpecializationChartResponseDTO {
			return dto.SpecializationChartResponseDTO{
				SpecializationName: item.SpecializationName,
				ContractCount:      item.ContractCount,
			}
		}),
	}, nil
}

func (c *controller) GetTemplatesChart(ctx context.Context, user models.User) (dto.TemplatesChartResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.TemplatesChartResponseDTO{}, err
	}
	templatesChart, err := c.repository.GetTemplatesChart(ctx)
	if err != nil {
		return dto.TemplatesChartResponseDTO{}, err
	}
	return dto.TemplatesChartResponseDTO{
		List: uslices.MapFunc(templatesChart, func(item entities.TemplateNameCountChart) dto.TemplateChartResponseDTO {
			return dto.TemplateChartResponseDTO{
				TemplateName:  item.TemplateName,
				ContractCount: item.ContractCount,
			}
		}),
	}, nil
}
