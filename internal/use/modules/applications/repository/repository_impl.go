package repository

import (
	"context"
	"encoding/json"

	"dogovorsBackEnd/internal/use/modules/applications/entities"
	"dogovorsBackEnd/internal/use/utils/scanners"
)

func (r *repository) CreateApplication(ctx context.Context, application entities.Application) (int, error) {
	query := `
INSERT INTO tbl_applications(educational_establishment_id, 
                             specialization_id, 
                             template_id, 
                             last_name, 
                             name, 
                             middle_name, 
                             phone_number, 
                             template_data)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING application_id
`
	templateDataJSON, err := json.Marshal(application.TemplateData)
	if err != nil {
		return 0, err
	}
	row := r.db.QueryRowContext(ctx, query,
		application.EducationalEstablishmentID,
		application.SpecializationID,
		application.TemplateID,
		application.LastName,
		application.Name,
		application.MiddleName,
		application.PhoneNumber,
		templateDataJSON)
	return scanners.Id(row)
}

func (r *repository) GetAllApplications(ctx context.Context) ([]entities.Application, error) {
	query := `
SELECT application_id,
       specialization_id, 
       educational_establishment_id, 
       template_id, 
       last_name,
       name,
       middle_name,
       phone_number,
       template_data,
       application_status from tbl_applications
`
	res, err := r.db.QueryContext(ctx, query)
	var applications []entities.Application
	for res.Next() {
		application, err := r.scanApplications(res)
		if err != nil {
			return nil, err
		}
		applications = append(applications, application)
	}
	return applications, err
}
