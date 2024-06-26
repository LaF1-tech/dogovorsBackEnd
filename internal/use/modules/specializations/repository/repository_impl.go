package repository

import (
	"context"
	"dogovorsBackEnd/internal/use/modules/specializations/entities"
)

func (r *repository) CreateSpecialization(ctx context.Context, specialization entities.Specialization) (id int, err error) {
	query := `
INSERT INTO tbl_specializations (specialization_name) VALUES ($1) RETURNING specialization_id;
`
	row := r.db.QueryRowContext(ctx, query, specialization.SpecializationName)
	err = row.Scan(&id)
	if err != nil {
		return 0, ErrCannotCreate
	}
	return id, nil
}

func (r *repository) GetSpecializations(ctx context.Context) ([]entities.Specialization, error) {
	query := `
SELECT specialization_id, specialization_name from tbl_specializations order by specialization_id
`
	res, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrNotFound
	}
	var specializations []entities.Specialization
	for res.Next() {
		preview, err := r.scanSpecializations(res)
		if err != nil {
			return nil, ErrNotFound
		}
		specializations = append(specializations, preview)
	}
	return specializations, nil
}

func (r *repository) GetSpecializationByID(ctx context.Context, id int) (entities.Specialization, error) {
	query := `
SELECT specialization_id, specialization_name from tbl_specializations WHERE specialization_id = $1
`
	response := r.db.QueryRowContext(ctx, query, id)
	return r.scanSpecialization(response)
}

func (r *repository) PatchSpecializationByID(ctx context.Context, specialization entities.Specialization) error {
	query := `
UPDATE tbl_specializations
SET specialization_name=COALESCE(NULLIF($2, ''), specialization_name)
WHERE specialization_id=$1
`
	_, err := r.db.ExecContext(ctx, query, specialization.SpecializationID, specialization.SpecializationName)
	if err != nil {
		return ErrCannotUpdate
	}
	return nil
}

func (r *repository) DeleteSpecializationByID(ctx context.Context, specID int) error {
	query := `
DELETE FROM tbl_specializations WHERE specialization_id = $1
`
	_, err := r.db.ExecContext(ctx, query, specID)
	if err != nil {
		return ErrCannotDelete
	}
	return nil
}
