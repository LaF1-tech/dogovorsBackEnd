package repository

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/templates/entities"
)

func (r *repository) GetAllTemplatesPreview(ctx context.Context) ([]entities.TemplatePreview, error) {
	query := `
SELECT template_id, template_name from tbl_templates
`
	res, err := r.db.QueryContext(ctx, query)
	var templates []entities.TemplatePreview
	for res.Next() {
		preview, err := r.scanTemplatePreview(res)
		if err != nil {
			return nil, err
		}
		templates = append(templates, preview)
	}
	return templates, err
}

func (r *repository) GetTemplateByID(ctx context.Context, id int) (entities.Template, error) {
	query := `
SELECT template_id, template_name, template_content, template_styles, necessary_data from tbl_templates WHERE template_id = $1
`
	res := r.db.QueryRowContext(ctx, query, id)
	return r.scanTemplate(res)
}

func (r *repository) GetFullTemplateDataByContractID(ctx context.Context, contractID int) (entities.FullTemplate, error) {
	query := `
SELECT contract_id, template_name, template_content, template_styles, necessary_data, data FROM vw_full_templates_data WHERE contract_id = $1
`
	res := r.db.QueryRowContext(ctx, query, contractID)
	return r.scanFullTemplate(res)
}
