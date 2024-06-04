package repository

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/charts/entities"
)

func (r *repository) GetPeriodChart(ctx context.Context) ([]entities.PeriodChart, error) {
	query := `SELECT DATE_TRUNC('month', execution_date) AS period,
       				 COUNT(*)                            AS contract_count
FROM tbl_contracts
GROUP BY period
ORDER BY period;
`
	response, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var data []entities.PeriodChart
	for response.Next() {
		periodChartData, err := r.scanPeriodChart(response)
		if err != nil {
			return nil, err
		}
		data = append(data, periodChartData)
	}
	return data, nil

}

func (r *repository) GetEducationalEstablishmentChart(ctx context.Context) ([]entities.EducationalEstablishmentCountChart, error) {
	query := `SELECT e.educational_establishment_name,
       COUNT(c.contract_id) AS contract_count
FROM tbl_contracts c
         JOIN
     tbl_students s ON c.student_id = s.student_id
         JOIN
     tbl_educational_establishments e ON s.educational_establishment_id = e.educational_establishment_id
GROUP BY e.educational_establishment_name
ORDER BY contract_count DESC;
`
	response, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var data []entities.EducationalEstablishmentCountChart
	for response.Next() {
		eduEstabChartData, err := r.scanEducationalEstablishmentChart(response)
		if err != nil {
			return nil, err
		}
		data = append(data, eduEstabChartData)
	}
	return data, nil

}

func (r *repository) GetSpecializationsChart(ctx context.Context) ([]entities.SpecializationNameCountChart, error) {
	query := `SELECT sp.specialization_name,
       COUNT(c.contract_id) AS contract_count
FROM tbl_contracts c
         JOIN
     tbl_students s ON c.student_id = s.student_id
         JOIN
     tbl_specializations sp ON s.specialization_id = sp.specialization_id
GROUP BY sp.specialization_name
ORDER BY contract_count DESC;`
	response, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var data []entities.SpecializationNameCountChart
	for response.Next() {
		specializationsChartData, err := r.scanSpecializationsChart(response)
		if err != nil {
			return nil, err
		}
		data = append(data, specializationsChartData)
	}
	return data, nil
}

func (r *repository) GetTemplatesChart(ctx context.Context) ([]entities.TemplateNameCountChart, error) {
	query := `SELECT t.template_name,
       COUNT(c.contract_id) AS contract_count
FROM tbl_contracts c
         JOIN
     tbl_templates t ON c.template_id = t.template_id
GROUP BY t.template_name
ORDER BY contract_count DESC;`
	response, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var data []entities.TemplateNameCountChart
	for response.Next() {
		templateChartData, err := r.scanTemplatesChart(response)
		if err != nil {
			return nil, err
		}
		data = append(data, templateChartData)
	}
	return data, nil
}
