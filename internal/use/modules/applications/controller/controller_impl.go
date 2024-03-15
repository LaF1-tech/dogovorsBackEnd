package controller

import (
	"context"

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
		TemplateID:                 request.TemplateID,
		LastName:                   request.LastName,
		Name:                       request.Name,
		MiddleName:                 request.MiddleName,
		PhoneNumber:                request.PhoneNumber,
		TemplateData:               request.TemplateData,
		ApplicationStatus:          request.ApplicationStatus,
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
		TemplateID:                 application.TemplateID,
		LastName:                   application.LastName,
		Name:                       application.Name,
		MiddleName:                 application.MiddleName,
		PhoneNumber:                application.PhoneNumber,
		TemplateData:               application.TemplateData,
		ApplicationStatus:          application.ApplicationStatus,
	}, nil
}

func (c *controller) GetAllApplications(ctx context.Context) (dto.ApplicationsResponseDTO, error) {
	applications, err := c.repository.GetAllApplications(ctx)
	if err != nil {
		return dto.ApplicationsResponseDTO{}, err
	}

	return dto.ApplicationsResponseDTO{
		List: uslices.MapFunc(applications, func(item entities.Application) dto.ApplicationItemResponseDTO {
			return dto.ApplicationItemResponseDTO{
				ApplicationID:              item.ApplicationID,
				EducationalEstablishmentID: item.EducationalEstablishmentID,
				SpecializationID:           item.SpecializationID,
				TemplateID:                 item.TemplateID,
				LastName:                   item.LastName,
				Name:                       item.Name,
				MiddleName:                 item.MiddleName,
				PhoneNumber:                item.PhoneNumber,
				TemplateData:               item.TemplateData,
				ApplicationStatus:          item.ApplicationStatus,
			}
		}),
	}, nil
}
