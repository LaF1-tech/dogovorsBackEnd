package repository

import (
	"database/sql"
	"dogovorsBackEnd/internal/use/modules/charts/entities"
)

func (r *repository) scanPeriodChart(row *sql.Rows) (chart entities.PeriodChart, err error) {
	err = row.Scan(
		&chart.Period,
		&chart.ContractCount,
	)
	if err != nil {
		return entities.PeriodChart{}, err
	}

	return
}

func (r *repository) scanEducationalEstablishmentChart(row *sql.Rows) (chart entities.EducationalEstablishmentCountChart, err error) {
	err = row.Scan(
		&chart.EducationalEstablishmentName,
		&chart.ContractCount,
	)
	if err != nil {
		return entities.EducationalEstablishmentCountChart{}, err
	}

	return
}

func (r *repository) scanSpecializationsChart(row *sql.Rows) (chart entities.SpecializationNameCountChart, err error) {
	err = row.Scan(
		&chart.SpecializationName,
		&chart.ContractCount,
	)
	if err != nil {
		return entities.SpecializationNameCountChart{}, err
	}

	return
}

func (r *repository) scanTemplatesChart(row *sql.Rows) (chart entities.TemplateNameCountChart, err error) {
	err = row.Scan(
		&chart.TemplateName,
		&chart.ContractCount,
	)
	if err != nil {
		return entities.TemplateNameCountChart{}, err
	}

	return
}
