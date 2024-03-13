package repository

import (
	"database/sql"
	"github.com/lib/pq/hstore"

	"dogovorsBackEnd/internal/use/modules/templates/entities"
)

func (r *repository) scanTemplatePreview(row *sql.Rows) (preview entities.TemplatePreview, err error) {
	err = row.Scan(&preview.TemplateID, &preview.TemplateName)
	return
}

func (r *repository) scanTemplate(row *sql.Row) (template entities.Template, err error) {
	var data hstore.Hstore

	err = row.Scan(
		&template.TemplateID,
		&template.TemplateName,
		&template.TemplateText,
		&data,
		&template.CreatedAt,
		&template.UpdatedAt,
	)

	template.NecessaryData = make(map[string]entities.TemplateVarType)

	for k, v := range data.Map {
		template.NecessaryData[k] = entities.TemplateVarType(v.String)
	}

	return
}

func (r *repository) scanFullTemplate(row *sql.Row) (fullTemplate entities.FullTemplate, err error) {
	var necessaryData hstore.Hstore

	err = row.Scan(
		&fullTemplate.ContractID,
		&fullTemplate.TemplateName,
		&fullTemplate.TemplateText,
		&necessaryData,
		&fullTemplate.Data,
	)

	for k, v := range necessaryData.Map {
		fullTemplate.NecessaryData[k] = entities.TemplateVarType(v.String)
	}

	return
}
