package controller

import (
	"context"
	"dogovorsBackEnd/internal/use/models"

	"dogovorsBackEnd/internal/use/modules/educationalestablishments/dto"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/entities"
	"dogovorsBackEnd/internal/use/utils/uslices"
)

func (c *controller) CreateEducationalEstablishment(ctx context.Context, user models.User, request dto.CreateEducationalEstablishmentRequestDTO) (dto.EducationalEstablishmentItemResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.EducationalEstablishmentItemResponseDTO{}, err
	}

	if err := c.validator.CreateEducationalEstablishmentRequest(ctx, request); err != nil {
		return dto.EducationalEstablishmentItemResponseDTO{}, err
	}

	e := entities.EducationalEstablishment{
		EducationalEstablishmentName:         request.EducationalEstablishmentName,
		EducationalEstablishmentContactPhone: request.EducationalEstablishmentContactPhone,
	}

	id, err := c.repository.CreateEducationalEstablishment(ctx, e)
	if err != nil {
		return dto.EducationalEstablishmentItemResponseDTO{}, err
	}

	e.EducationalEstablishmentID = id

	return dto.EducationalEstablishmentItemResponseDTO{
		EducationalEstablishmentID:           e.EducationalEstablishmentID,
		EducationalEstablishmentName:         e.EducationalEstablishmentName,
		EducationalEstablishmentContactPhone: e.EducationalEstablishmentContactPhone,
	}, nil
}

func (c *controller) GetEducationalEstablishments(ctx context.Context) (dto.EducationalEstablishmentsResponseDTO, error) {
	educationalEstablishments, err := c.repository.GetEducationalEstablishments(ctx)
	if err != nil {
		return dto.EducationalEstablishmentsResponseDTO{}, err
	}

	return dto.EducationalEstablishmentsResponseDTO{
		List: uslices.MapFunc(educationalEstablishments, func(item entities.EducationalEstablishment) dto.EducationalEstablishmentItemResponseDTO {
			return dto.EducationalEstablishmentItemResponseDTO{
				EducationalEstablishmentID:           item.EducationalEstablishmentID,
				EducationalEstablishmentName:         item.EducationalEstablishmentName,
				EducationalEstablishmentContactPhone: item.EducationalEstablishmentContactPhone,
			}
		}),
	}, nil
}

func (c *controller) GetEducationalEstablishmentByID(ctx context.Context, id int) (dto.EducationalEstablishmentItemResponseDTO, error) {
	establishment, err := c.repository.GetEducationalEstablishmentByID(ctx, id)
	if err != nil {
		return dto.EducationalEstablishmentItemResponseDTO{}, err
	}
	return dto.EducationalEstablishmentItemResponseDTO{
		EducationalEstablishmentID:           establishment.EducationalEstablishmentID,
		EducationalEstablishmentName:         establishment.EducationalEstablishmentName,
		EducationalEstablishmentContactPhone: establishment.EducationalEstablishmentContactPhone,
	}, nil
}

func (c *controller) PatchEducationalEstablishmentByID(ctx context.Context, user models.User, request dto.PatchEducationalEstablishmentRequestDTO) error {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return err
	}

	if err := c.validator.PatchEducationalEstablishmentRequest(ctx, request); err != nil {
		return err
	}

	establishment := entities.EducationalEstablishment{
		EducationalEstablishmentID:           request.EducationalEstablishmentID,
		EducationalEstablishmentName:         request.EducationalEstablishmentName,
		EducationalEstablishmentContactPhone: request.EducationalEstablishmentContactPhone,
	}

	err := c.repository.PatchEducationalEstablishmentByID(ctx, establishment)
	if err != nil {
		return err
	}
	return nil
}

func (c *controller) DeleteEducationalEstablishmentByID(ctx context.Context, user models.User, id int) error {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return err
	}
	return c.repository.DeleteEducationalEstablishmentByID(ctx, id)
}
