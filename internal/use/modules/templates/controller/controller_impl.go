package controller

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/templates/dto"
	"dogovorsBackEnd/internal/use/modules/templates/entities"
	"dogovorsBackEnd/internal/use/utils/uslices"
)

func (c *controller) GetTemplatesPreview(ctx context.Context) (dto.TemplatePreviewResponseDTO, error) {
	templatesPreview, err := c.repository.GetAllTemplatesPreview(ctx)
	if err != nil {
		return dto.TemplatePreviewResponseDTO{}, err
	}

	return dto.TemplatePreviewResponseDTO{
		List: uslices.MapFunc(templatesPreview, func(item entities.TemplatePreview) dto.TemplatePreviewItemResponseDTO {
			return dto.TemplatePreviewItemResponseDTO{
				TemplateID:   item.TemplateID,
				TemplateName: item.TemplateName,
			}
		}),
	}, nil
}

func (c *controller) GetTemplateByID(ctx context.Context, id int) (dto.TemplateResponseDTO, error) {
	template, err := c.repository.GetTemplateByID(ctx, id)
	if err != nil {
		return dto.TemplateResponseDTO{}, err
	}

	return dto.TemplateResponseDTO{
		TemplateID:    template.TemplateID,
		TemplateName:  template.TemplateName,
		NecessaryData: template.NecessaryData,
	}, nil
}

func (c *controller) GenerateTemplatePDF(ctx context.Context, id int) error {

	return nil
}
