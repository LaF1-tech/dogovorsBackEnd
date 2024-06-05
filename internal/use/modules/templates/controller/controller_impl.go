package controller

import (
	"context"
	"dogovorsBackEnd/internal/use/models"

	"dogovorsBackEnd/internal/use/modules/templates/dto"
	"dogovorsBackEnd/internal/use/modules/templates/entities"
	"dogovorsBackEnd/internal/use/utils/uslices"
)

func (c *controller) CreateTemplate(ctx context.Context, user models.User, request dto.CreateTemplateRequestDTO) (dto.TemplateResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.TemplateResponseDTO{}, err
	}

	template := entities.Template{
		TemplateName:    request.TemplateName,
		TemplateContent: request.TemplateContent,
		TemplateStyles:  request.TemplateStyles,
		NecessaryData:   request.NecessaryData,
	}

	id, err := c.repository.CreateTemplate(ctx, template)
	if err != nil {
		return dto.TemplateResponseDTO{}, err
	}

	template.TemplateID = id

	return dto.TemplateResponseDTO{
		TemplateID:    template.TemplateID,
		TemplateName:  template.TemplateName,
		NecessaryData: template.NecessaryData,
	}, nil
}

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
		TemplateID:      template.TemplateID,
		TemplateName:    template.TemplateName,
		TemplateContent: template.TemplateContent,
		TemplateStyles:  template.TemplateStyles,
		NecessaryData:   template.NecessaryData,
	}, nil
}

func (c *controller) DeleteTemplateByID(ctx context.Context, user models.User, id int) error {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return err
	}
	return c.repository.DeleteTemplateByID(ctx, id)
}

func (c *controller) PatchTemplateByID(ctx context.Context, user models.User, request dto.PatchTemplateRequestDTO) error {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return err
	}

	template := entities.Template{
		TemplateID:      request.TemplateID,
		TemplateName:    request.TemplateName,
		TemplateContent: request.TemplateContent,
		TemplateStyles:  request.TemplateStyles,
		NecessaryData:   request.NecessaryData,
	}

	err := c.repository.PatchTemplateByID(ctx, template)
	if err != nil {
		return err
	}

	return nil
}

func (c *controller) GenerateTemplatePDF(ctx context.Context, id int) error {

	return nil
}
