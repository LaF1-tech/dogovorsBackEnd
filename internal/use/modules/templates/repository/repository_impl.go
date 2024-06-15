package repository

import (
	"context"
	"database/sql"
	"dogovorsBackEnd/internal/use/modules/templates/entities"
	"dogovorsBackEnd/internal/use/utils/scanners"
	"github.com/lib/pq/hstore"
)

func (r *repository) CreateTemplate(ctx context.Context, template entities.Template) (int, error) {
	query := `
INSERT INTO tbl_templates (template_name, template_content, template_styles, necessary_data) VALUES ($1,$2,$3,$4) RETURNING template_id
`

	var hueta hstore.Hstore
	hueta.Map = make(map[string]sql.NullString)

	for k, v := range template.NecessaryData {
		hueta.Map[k] = sql.NullString{String: string(v), Valid: true}
	}

	row := r.db.QueryRowContext(ctx, query, template.TemplateName, template.TemplateContent, template.TemplateStyles, hueta)
	return scanners.Id(row)
}

func (r *repository) GetAllTemplatesPreview(ctx context.Context) ([]entities.TemplatePreview, error) {
	query := `
SELECT template_id, template_name from tbl_templates order by template_id
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

func (r *repository) PatchTemplateByID(ctx context.Context, template entities.Template) error {
	query := `
UPDATE tbl_templates 
SET template_name=COALESCE(NULLIF($2, ''), template_name),
    template_content=COALESCE(NULLIF($3, ''::text), template_content), 
    template_styles=COALESCE(NULLIF($4, ''::text), template_styles), 
    necessary_data=COALESCE(NULLIF($5::hstore, '')::hstore, necessary_data)
WHERE template_id=$1
`
	var hueta hstore.Hstore
	hueta.Map = make(map[string]sql.NullString)

	for k, v := range template.NecessaryData {
		hueta.Map[k] = sql.NullString{String: string(v), Valid: true}
	}

	_, err := r.db.ExecContext(ctx, query,
		template.TemplateID,
		template.TemplateName,
		template.TemplateContent,
		template.TemplateStyles,
		hueta,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteTemplateByID(ctx context.Context, templateID int) error {
	query := `
DELETE FROM tbl_templates WHERE template_id = $1
`
	_, err := r.db.ExecContext(ctx, query, templateID)
	if err != nil {
		return err
	}
	return nil
}
