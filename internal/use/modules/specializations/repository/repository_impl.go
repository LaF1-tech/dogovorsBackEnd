package repository

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/specializations/entities"
)

func (r *repository) GetEducationalEstablishments(ctx context.Context) ([]entities.Specialization, error) {
	query := `
SELECT specialization_id, specialization_name from tbl_specializations
`
	res, err := r.db.QueryContext(ctx, query)
	var specializations []entities.Specialization
	for res.Next() {
		preview, err := r.scanSpecialization(res)
		if err != nil {
			return nil, err
		}
		specializations = append(specializations, preview)
	}
	return specializations, err
}
