package repository

import (
	"database/sql"
	"encoding/json"

	"dogovorsBackEnd/internal/use/modules/applications/entities"
)

func (r *repository) scanAggregatedApplication(row *sql.Rows) (application entities.AggregatedApplication, err error) {
	var data []uint8

	err = row.Scan(
		&application.ApplicationID,
		&application.SpecializationName,
		&application.EducationalEstablishmentName,
		&application.LastName,
		&application.Name,
		&application.MiddleName,
		&application.PhoneNumber,
		&data,
		&application.ApplicationStatus,
		&application.ExecutionDate,
		&application.ExpirationDate)
	if err != nil {
		return entities.AggregatedApplication{}, err
	}

	err = json.Unmarshal(data, &application.Types)
	if err != nil {
		return entities.AggregatedApplication{}, err
	}

	return
}

func (r *repository) scanApplication(row *sql.Row) (application entities.Application, err error) {
	var data []uint8

	err = row.Scan(
		&application.ApplicationID,
		&application.EducationalEstablishmentID,
		&application.SpecializationID,
		&application.LastName,
		&application.Name,
		&application.MiddleName,
		&application.PhoneNumber,
		&data,
		&application.ApplicationStatus,
		&application.ExecutionDate,
		&application.ExpirationDate)
	if err != nil {
		return entities.Application{}, err
	}

	err = json.Unmarshal(data, &application.Types)
	if err != nil {
		return entities.Application{}, err
	}

	return
}
