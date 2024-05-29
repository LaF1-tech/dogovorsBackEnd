package repository

import (
	"context"
	"encoding/json"

	"dogovorsBackEnd/internal/use/modules/applications/entities"
	contracts "dogovorsBackEnd/internal/use/modules/contracts/entities"
	"dogovorsBackEnd/internal/use/utils/scanners"
)

func (r *repository) CreateApplication(ctx context.Context, application entities.Application) (int, error) {
	query := `
INSERT INTO tbl_applications(educational_establishment_id, 
                             specialization_id, 
                             last_name, 
                             name, 
                             middle_name, 
                             phone_number, 
                             types,
                             execution_date,
                             expiration_date)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8, $9) RETURNING application_id
`
	typesJSON, err := json.Marshal(application.Types)
	if err != nil {
		return 0, err
	}
	row := r.db.QueryRowContext(ctx, query,
		application.EducationalEstablishmentID,
		application.SpecializationID,
		application.LastName,
		application.Name,
		application.MiddleName,
		application.PhoneNumber,
		typesJSON,
		application.ExecutionDate,
		application.ExpirationDate)
	return scanners.Id(row)
}

func (r *repository) GetAllApplications(ctx context.Context) ([]entities.AggregatedApplication, error) {
	query := `
SELECT application_id,
       specialization_name, 
       educational_establishment_name, 
       last_name,
       name,
       middle_name,
       phone_number,
       types,
       application_status,
       execution_date,
       expiration_date from vw_full_application_data
`
	res, err := r.db.QueryContext(ctx, query)
	var applications []entities.AggregatedApplication
	for res.Next() {
		application, err := r.scanAggregatedApplication(res)
		if err != nil {
			return nil, err
		}
		applications = append(applications, application)
	}
	return applications, err
}

func (r *repository) GetApplicationByID(ctx context.Context, id int) (entities.Application, error) {
	query := `
SELECT application_id, 
       educational_establishment_id, 
       specialization_id,
       last_name,
       name,
       middle_name,
       phone_number, 
       types,
       application_status,
       execution_date,
       expiration_date
from tbl_applications WHERE application_id=$1
`
	response := r.db.QueryRowContext(ctx, query, id)
	return r.scanApplication(response)
}

func (r *repository) PatchApplicationByID(ctx context.Context, application entities.Application) error {
	query := `UPDATE tbl_applications
SET educational_establishment_id=COALESCE(NULLIF($1, 0), educational_establishment_id),
    specialization_id=COALESCE(NULLIF($2, 0), specialization_id),
    last_name=COALESCE(NULLIF($3, ''), last_name),
    name=COALESCE(NULLIF($4, ''), name),
    middle_name=COALESCE(NULLIF($5, ''), middle_name),
    phone_number=COALESCE(NULLIF($6, ''), phone_number),
    types=COALESCE(NULLIF($7::text, 'null')::json, types),
    application_status=COALESCE(NULLIF($8, '')::tp_application_status, application_status)
WHERE application_id = $9
`
	typesJSON, err := json.Marshal(application.Types)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query,
		application.EducationalEstablishmentID,
		application.SpecializationID,
		application.LastName,
		application.Name,
		application.MiddleName,
		application.PhoneNumber,
		typesJSON,
		application.ApplicationStatus,
		application.ApplicationID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateContract(ctx context.Context, contract contracts.Contract) (int, error) {
	query := `
INSERT INTO tbl_contracts (student_id, employee_id, template_id, execution_date, expiration_date) VALUES ($1,$2,$3,$4,$5) RETURNING contract_id
`

	row := r.db.QueryRowContext(ctx, query, contract.StudentID, contract.EmployeeID, contract.TemplateId, contract.ExecutionDate, contract.ExpirationDate)
	return scanners.Id(row)
}

func (r *repository) CreateStudent(ctx context.Context, student entities.Student) (int, error) {
	query := `
INSERT INTO tbl_students (student_name, student_middle_name, student_last_name, specialization_id, educational_establishment_id, student_phone_number) VALUES ($1,$2,$3,$4,$5,$6) RETURNING student_id
`

	row := r.db.QueryRowContext(ctx, query, student.StudentName, student.StudentMiddleName, student.StudentLastName, student.SpecializationID, student.EducationalEstablishmentID, student.StudentPhoneNumber)
	return scanners.Id(row)
}

func (r *repository) CreateContractDataRegistryEntry(ctx context.Context, cdr entities.ContractDataRegistry) (int, error) {
	query := `
INSERT INTO tbl_contract_data_registry (template_id, contract_id, data) VALUES ($1,$2,$3) RETURNING contract_id
`
	contractData, err := json.Marshal(cdr.Data)
	if err != nil {
		return 0, err
	}
	row := r.db.QueryRowContext(ctx, query, cdr.TemplateID, cdr.ContractID, contractData)
	return scanners.Id(row)
}
