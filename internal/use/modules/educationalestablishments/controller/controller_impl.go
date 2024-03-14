package controller

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/educationalestablishments/dto"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/entities"
	"dogovorsBackEnd/internal/use/utils/uslices"
)

func (c *controller) GetEducationalEstablishments(ctx context.Context) (dto.EducationalEstablishmentsResponseDTO, error) {
	educationalEstablishments, err := c.repository.GetEducationalEstablishments(ctx)
	if err != nil {
		return dto.EducationalEstablishmentsResponseDTO{}, err
	}

	return dto.EducationalEstablishmentsResponseDTO{
		List: uslices.MapFunc(educationalEstablishments, func(item entities.EducationalEstablishments) dto.EducationalEstablishmentsItemResponseDTO {
			return dto.EducationalEstablishmentsItemResponseDTO{
				EducationalEstablishmentID:           item.EducationalEstablishmentID,
				EducationalEstablishmentName:         item.EducationalEstablishmentName,
				EducationalEstablishmentContactPhone: item.EducationalEstablishmentContactPhone,
			}
		}),
	}, nil
}
