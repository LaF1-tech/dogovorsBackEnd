package repository

import (
	"database/sql"
	"encoding/json"
	"github.com/lib/pq/hstore"

	"dogovorsBackEnd/internal/use/modules/contracts/entities"
	tEntities "dogovorsBackEnd/internal/use/modules/templates/entities"
)

func (r *repository) scanAggregatedContract(row *sql.Rows) (contract entities.AggregatedContract, err error) {
	err = row.Scan(
		&contract.ContractID,
		&contract.StudentName,
		&contract.StudentLastName,
		&contract.EmployeeFirstName,
		&contract.EmployeeLastName,
		&contract.TemplateName,
		&contract.ExecutionDate,
		&contract.ExpirationDate)
	if err != nil {
		return entities.AggregatedContract{}, err
	}
	return
}

func (r *repository) scanContract(row *sql.Row) (contract entities.Contract, err error) {

	err = row.Scan(
		&contract.ContractID,
		&contract.StudentID,
		&contract.EmployeeID,
		&contract.TemplateId,
		&contract.ExecutionDate,
		&contract.ExpirationDate)
	if err != nil {
		return entities.Contract{}, err
	}

	return
}

func (r *repository) scanTemplate(row *sql.Row) (template tEntities.Template, err error) {
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
		return tEntities.Template{}, err
	}

	if templateStyles == nil {
		template.TemplateStyles = ""
	} else {
		template.TemplateStyles = *templateStyles
	}

	template.NecessaryData = make(map[string]tEntities.TemplateVarType)

	for k, v := range data.Map {
		template.NecessaryData[k] = tEntities.TemplateVarType(v.String)
	}

	return
}

func (r *repository) scanFullTemplate(row *sql.Row) (fullTemplate tEntities.FullTemplate, err error) {
	var necessaryData hstore.Hstore
	var data []byte
	var templateStyles *string

	err = row.Scan(
		&fullTemplate.ContractID,
		&fullTemplate.TemplateName,
		&fullTemplate.TemplateContent,
		&templateStyles,
		&necessaryData,
		&data,
	)
	if err != nil {
		return tEntities.FullTemplate{}, err
	}

	if templateStyles == nil {
		fullTemplate.TemplateStyles = ""
	} else {
		fullTemplate.TemplateStyles = *templateStyles
	}

	err = json.Unmarshal(data, &fullTemplate.Data)
	if err != nil {
		return tEntities.FullTemplate{}, err
	}

	fullTemplate.NecessaryData = make(map[string]tEntities.TemplateVarType)
	for k, v := range necessaryData.Map {
		fullTemplate.NecessaryData[k] = tEntities.TemplateVarType(v.String)
	}

	return
}
