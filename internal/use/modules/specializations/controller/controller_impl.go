package controller

import (
	"context"
	"dogovorsBackEnd/internal/use/models"

	"dogovorsBackEnd/internal/use/modules/specializations/dto"
	"dogovorsBackEnd/internal/use/modules/specializations/entities"
	"dogovorsBackEnd/internal/use/utils/uslices"
)

func (c *controller) CreateSpecialization(ctx context.Context, user models.User, request dto.CreateSpecializationRequestDTO) (dto.SpecializationItemResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.SpecializationItemResponseDTO{}, err
	}

	if err := c.validator.CreateSpecializationRequest(ctx, request); err != nil {
		return dto.SpecializationItemResponseDTO{}, err
	}

	s := entities.Specialization{
		SpecializationName: request.SpecializationName,
	}

	id, err := c.repository.CreateSpecialization(ctx, s)
	if err != nil {
		return dto.SpecializationItemResponseDTO{}, err
	}

	s.SpecializationID = id

	return dto.SpecializationItemResponseDTO{
		SpecializationID:   s.SpecializationID,
		SpecializationName: s.SpecializationName,
	}, nil
}

func (c *controller) GetSpecializations(ctx context.Context) (dto.SpecializationsResponseDTO, error) {
	educationalEstablishments, err := c.repository.GetSpecializations(ctx)
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

func (c *controller) GetSpecializationByID(ctx context.Context, id int) (dto.SpecializationItemResponseDTO, error) {
	specialization, err := c.repository.GetSpecializationByID(ctx, id)
	if err != nil {
		return dto.SpecializationItemResponseDTO{}, err
	}
	return dto.SpecializationItemResponseDTO{
		SpecializationID:   specialization.SpecializationID,
		SpecializationName: specialization.SpecializationName,
	}, nil
}

func (c *controller) PatchSpecializationByID(ctx context.Context, user models.User, request dto.PatchSpecializationRequestDTO) error {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return err
	}

	if err := c.validator.PatchSpecializationRequest(ctx, request); err != nil {
		return err
	}

	specialization := entities.Specialization{
		SpecializationID:   request.SpecializationID,
		SpecializationName: request.SpecializationName,
	}

	err := c.repository.PatchSpecializationByID(ctx, specialization)
	if err != nil {
		return err
	}
	return nil
}

func (c *controller) DeleteSpecializationByID(ctx context.Context, user models.User, id int) error {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return err
	}
	return c.repository.DeleteSpecializationByID(ctx, id)
}
