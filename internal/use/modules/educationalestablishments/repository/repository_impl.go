package repository

import (
	"context"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/entities"
)

func (r *repository) CreateEducationalEstablishment(ctx context.Context, establishment entities.EducationalEstablishment) (id int, err error) {
	query := `
INSERT INTO tbl_educational_establishments (educational_establishment_name, educational_establishment_contact_phone) VALUES ($1,$2) RETURNING educational_establishment_id
`

	row := r.db.QueryRowContext(ctx, query, establishment.EducationalEstablishmentName, establishment.EducationalEstablishmentContactPhone)

	err = row.Scan(&id)
	if err != nil {
		return 0, ErrCannotCreate
	}
	return id, nil
}

func (r *repository) GetEducationalEstablishments(ctx context.Context) ([]entities.EducationalEstablishment, error) {
	query := `
SELECT educational_establishment_id,
	   educational_establishment_name,
	   educational_establishment_contact_phone 
from tbl_educational_establishments order by educational_establishment_id
`
	res, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrNotFound
	}
	var educationalEstablishments []entities.EducationalEstablishment
	for res.Next() {
		preview, err := r.scanEducationalEstablishments(res)
		if err != nil {
			return nil, ErrNotFound
		}
		educationalEstablishments = append(educationalEstablishments, preview)
	}
	return educationalEstablishments, nil
}

func (r *repository) GetEducationalEstablishmentByID(ctx context.Context, id int) (entities.EducationalEstablishment, error) {
	query := `
SELECT educational_establishment_id,
	   educational_establishment_name,
	   educational_establishment_contact_phone 
from tbl_educational_establishments WHERE educational_establishment_id = $1
`
	response := r.db.QueryRowContext(ctx, query, id)
	return r.scanEducationalEstablishment(response)
}

func (r *repository) PatchEducationalEstablishmentByID(ctx context.Context, establishment entities.EducationalEstablishment) error {
	query := `
UPDATE tbl_educational_establishments
SET educational_establishment_name=COALESCE(NULLIF($2, ''), educational_establishment_name),
	educational_establishment_contact_phone=COALESCE(NULLIF($3, ''), educational_establishment_contact_phone)
WHERE educational_establishment_id=$1
`

	_, err := r.db.ExecContext(ctx, query,
		establishment.EducationalEstablishmentID,
		establishment.EducationalEstablishmentName,
		establishment.EducationalEstablishmentContactPhone,
	)

	if err != nil {
		return ErrCannotUpdate
	}
	return nil
}

func (r *repository) DeleteEducationalEstablishmentByID(ctx context.Context, id int) error {
	query := `
DELETE FROM tbl_educational_establishments WHERE educational_establishment_id=$1
`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return ErrCannotDelete
	}
	return nil
}
