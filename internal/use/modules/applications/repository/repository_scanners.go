package repository

import (
	"database/sql"
	"encoding/json"

	"dogovorsBackEnd/internal/use/modules/applications/entities"
)

func (r *repository) scanApplications(row *sql.Rows) (application entities.Application, err error) {
	var data []uint8

	err = row.Scan(
		&application.ApplicationID,
		&application.SpecializationID,
		&application.EducationalEstablishmentID,
		&application.TemplateID,
		&application.LastName,
		&application.Name,
		&application.MiddleName,
		&application.PhoneNumber,
		&data,
		&application.ApplicationStatus)
	if err != nil {
		return entities.Application{}, err
	}

	err = json.Unmarshal(data, &application.TemplateData)
	if err != nil {
		return entities.Application{}, err
	}

	return
}
