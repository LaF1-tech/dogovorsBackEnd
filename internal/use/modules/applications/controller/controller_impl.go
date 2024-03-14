package controller

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/applications/dto"
	"dogovorsBackEnd/internal/use/modules/applications/entities"
)

func (c *controller) CreateApplication(ctx context.Context, request dto.CreateApplicationRequestDTO) (dto.CreateApplicationResponseDTO, error) {
	if err := c.validator.CreateApplicationRequest(ctx, request); err != nil {
		return dto.CreateApplicationResponseDTO{}, err
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
		return dto.CreateApplicationResponseDTO{}, err
	}

	application.ApplicationID = id

	return dto.CreateApplicationResponseDTO{
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
