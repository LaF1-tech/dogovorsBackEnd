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
	var templateStyles *string

	err = row.Scan(
		&template.TemplateID,
		&template.TemplateName,
		&template.TemplateContent,
		&templateStyles,
		&data,
	)
	if err != nil {
		return entities.Template{}, err
	}

	if templateStyles == nil {
		template.TemplateStyles = ""
	} else {
		template.TemplateStyles = *templateStyles
	}

	template.NecessaryData = make(map[string]entities.TemplateVarType)

	for k, v := range data.Map {
		template.NecessaryData[k] = entities.TemplateVarType(v.String)
	}

	return
}

func (r *repository) scanFullTemplate(row *sql.Row) (fullTemplate entities.FullTemplate, err error) {
	var necessaryData hstore.Hstore
	var templateStyles *string

	err = row.Scan(
		&fullTemplate.ContractID,
		&fullTemplate.TemplateName,
		&fullTemplate.TemplateContent,
		&templateStyles,
		&necessaryData,
		&fullTemplate.Data,
	)
	if err != nil {
		return entities.FullTemplate{}, err
	}

	if templateStyles == nil {
		fullTemplate.TemplateStyles = ""
	} else {
		fullTemplate.TemplateStyles = *templateStyles
	}

	for k, v := range necessaryData.Map {
		fullTemplate.NecessaryData[k] = entities.TemplateVarType(v.String)
	}

	return
}
