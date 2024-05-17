package controller

import (
	"context"
	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/controller/validator"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/dto"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/repository"
)

type EducationalEstablishmentsController interface {
	CreateEducationalEstablishment(ctx context.Context, user models.User, request dto.CreateEducationalEstablishmentRequestDTO) (dto.EducationalEstablishmentItemResponseDTO, error)
	GetEducationalEstablishments(ctx context.Context) (dto.EducationalEstablishmentsResponseDTO, error)
	GetEducationalEstablishmentByID(ctx context.Context, id int) (dto.EducationalEstablishmentItemResponseDTO, error)
	PatchEducationalEstablishmentByID(ctx context.Context, user models.User, request dto.PatchEducationalEstablishmentRequestDTO) error
	DeleteEducationalEstablishmentByID(ctx context.Context, user models.User, id int) error
}

type controller struct {
	validator  validator.EducationalEstablishmentValidator
	repository repository.EducationalEstablishmentsRepository
}

func New(validator validator.EducationalEstablishmentValidator, repository repository.EducationalEstablishmentsRepository) EducationalEstablishmentsController {
	return &controller{validator: validator, repository: repository}
}
