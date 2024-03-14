package repository

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/educationalestablishments/entities"
)

func (r *repository) GetEducationalEstablishments(ctx context.Context) ([]entities.EducationalEstablishments, error) {
	query := `
SELECT educational_establishment_id, educational_establishment_name, educational_establishment_contact_phone from tbl_educational_establishments
`
	res, err := r.db.QueryContext(ctx, query)
	var educationalEstablishments []entities.EducationalEstablishments
	for res.Next() {
		preview, err := r.scanEducationalEstablishment(res)
		if err != nil {
			return nil, err
		}
		educationalEstablishments = append(educationalEstablishments, preview)
	}
	return educationalEstablishments, err
}
