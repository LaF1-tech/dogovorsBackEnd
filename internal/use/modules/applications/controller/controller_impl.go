package controller

import (
	"context"

	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/applications/dto"
	"dogovorsBackEnd/internal/use/modules/applications/entities"
	"dogovorsBackEnd/internal/use/utils/uslices"
)

func (c *controller) CreateApplication(ctx context.Context, request dto.CreateApplicationRequestDTO) (dto.ApplicationItemResponseDTO, error) {
	if err := c.validator.CreateApplicationRequest(ctx, request); err != nil {
		return dto.ApplicationItemResponseDTO{}, err
	}

	application := entities.Application{
		EducationalEstablishmentID: request.EducationalEstablishmentID,
		SpecializationID:           request.SpecializationID,
		LastName:                   request.LastName,
		Name:                       request.Name,
		MiddleName:                 request.MiddleName,
		PhoneNumber:                request.PhoneNumber,
		Types:                      request.Types,
		ApplicationStatus:          request.ApplicationStatus,
		ExecutionDate:              request.ExecutionDate,
		ExpirationDate:             request.ExpirationDate,
	}

	id, err := c.repository.CreateApplication(ctx, application)
	if err != nil {
		return dto.ApplicationItemResponseDTO{}, err
	}

	application.ApplicationID = id

	return dto.ApplicationItemResponseDTO{
		ApplicationID:              application.ApplicationID,
		EducationalEstablishmentID: application.EducationalEstablishmentID,
		SpecializationID:           application.SpecializationID,
		LastName:                   application.LastName,
		Name:                       application.Name,
		MiddleName:                 application.MiddleName,
		PhoneNumber:                application.PhoneNumber,
		Types:                      application.Types,
		ApplicationStatus:          application.ApplicationStatus,
		ExecutionDate:              application.ExecutionDate,
		ExpirationDate:             application.ExpirationDate,
	}, nil
}

func (c *controller) GetAllApplications(ctx context.Context, user models.User) (dto.ApplicationsViewResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.ApplicationsViewResponseDTO{}, err
	}

	applications, err := c.repository.GetAllApplications(ctx)
	if err != nil {
		return dto.ApplicationsViewResponseDTO{}, err
	}

	return dto.ApplicationsViewResponseDTO{
		List: uslices.MapFunc(applications, func(item entities.AggregatedApplication) dto.ApplicationViewResponseDTO {
			return dto.ApplicationViewResponseDTO{
				ApplicationID:                item.ApplicationID,
				EducationalEstablishmentName: item.EducationalEstablishmentName,
				SpecializationName:           item.SpecializationName,
				LastName:                     item.LastName,
				Name:                         item.Name,
				MiddleName:                   item.MiddleName,
				PhoneNumber:                  item.PhoneNumber,
				Types:                        item.Types,
				ApplicationStatus:            item.ApplicationStatus,
				ExecutionDate:                item.ExecutionDate,
				ExpirationDate:               item.ExpirationDate,
			}
		}),
	}, nil
}

func (c *controller) PatchApplicationByID(ctx context.Context, user models.User, request dto.PatchApplicationRequestDTO) error {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return err
	}

	application := entities.Application{
		ApplicationID:              request.ApplicationID,
		EducationalEstablishmentID: request.EducationalEstablishmentID,
		SpecializationID:           request.SpecializationID,
		LastName:                   request.LastName,
		Name:                       request.Name,
		MiddleName:                 request.MiddleName,
		PhoneNumber:                request.PhoneNumber,
		Types:                      request.Types,
		ApplicationStatus:          request.ApplicationStatus,
		ExecutionDate:              request.ExecutionDate,
		ExpirationDate:             request.ExpirationDate,
	}

	err := c.repository.PatchApplicationByID(ctx, application)
	if err != nil {
		return err
	}

	if application.ApplicationStatus == entities.ApplicationStatusAccepted {
		if err = c.acceptApplication(ctx, user, request); err != nil {
			return err
		}
	}

	return nil
}
