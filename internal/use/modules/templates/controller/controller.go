package controller

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/templates/controller/validator"
	"dogovorsBackEnd/internal/use/modules/templates/dto"
	"dogovorsBackEnd/internal/use/modules/templates/repository"
)

type TemplatesController interface {
	GetTemplatesPreview(ctx context.Context) (dto.TemplatePreviewResponseDTO, error)
	GetTemplateByID(ctx context.Context, id int) (dto.TemplateResponseDTO, error)
	GenerateTemplatePDF(ctx context.Context, id int) error
}

type controller struct {
	validator validator.TemplateValidator

	repository repository.TemplatesRepository
}

func New(validator validator.TemplateValidator, repository repository.TemplatesRepository) TemplatesController {
	return &controller{validator: validator, repository: repository}
}
