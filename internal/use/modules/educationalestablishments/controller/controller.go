package controller

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/educationalestablishments/controller/validator"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/dto"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/repository"
)

type EducationalEstablishmentsController interface {
	GetEducationalEstablishments(ctx context.Context) (dto.EducationalEstablishmentsResponseDTO, error)
}

type controller struct {
	validator  validator.EducationalEstablishmentValidator
	repository repository.EducationalEstablishmentsRepository
}

func New(validator validator.EducationalEstablishmentValidator, repository repository.EducationalEstablishmentsRepository) EducationalEstablishmentsController {
	return &controller{validator: validator, repository: repository}
}
