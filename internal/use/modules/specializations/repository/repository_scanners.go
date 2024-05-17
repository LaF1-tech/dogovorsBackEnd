package repository

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/specializations/entities"
)

func (r *repository) scanSpecializations(row *sql.Rows) (specialization entities.Specialization, err error) {
	err = row.Scan(&specialization.SpecializationID, &specialization.SpecializationName)
	return
}
func (r *repository) scanSpecialization(row *sql.Row) (specialization entities.Specialization, err error) {
	err = row.Scan(&specialization.SpecializationID, &specialization.SpecializationName)
	return
}
