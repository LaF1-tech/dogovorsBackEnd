package controller

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/specializations/dto"
	"dogovorsBackEnd/internal/use/modules/specializations/entities"
	"dogovorsBackEnd/internal/use/utils/uslices"
)

func (c *controller) GetSpecializations(ctx context.Context) (dto.SpecializationsResponseDTO, error) {
	educationalEstablishments, err := c.repository.GetEducationalEstablishments(ctx)
	if err != nil {
		return dto.SpecializationsResponseDTO{}, err
	}

	return dto.SpecializationsResponseDTO{
		List: uslices.MapFunc(educationalEstablishments, func(item entities.Specialization) dto.SpecializationItemResponseDTO {
			return dto.SpecializationItemResponseDTO{
				SpecializationID:   item.SpecializationID,
				SpecializationName: item.SpecializationName,
			}
		}),
	}, nil
}
