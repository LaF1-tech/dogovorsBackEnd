package repository

import (
	"bytes"
	"context"
	"html/template"

	tEntities "dogovorsBackEnd/internal/use/modules/templates/entities"
)

func (r *repository) GetAggregatedTemplateByContractID(ctx context.Context, id int) (tEntities.FullTemplate, error) {
	query := `
SELECT contract_id, template_name, template_content, template_styles, necessary_data, data from vw_full_templates_data WHERE contract_id = $1
`
	res := r.db.QueryRowContext(ctx, query, id)
	return r.scanFullTemplate(res)
}

func (r *repository) GenerateTemplateContent(_ context.Context, fullTemplate tEntities.FullTemplate) (string, error) {
	t, err := template.New("").Parse(fullTemplate.TemplateContent)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer
	err = t.Execute(&buff, fullTemplate.Data)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
