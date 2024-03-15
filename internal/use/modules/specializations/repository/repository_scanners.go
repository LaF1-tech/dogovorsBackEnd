package repository

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/specializations/entities"
)

func (r *repository) scanSpecialization(row *sql.Rows) (specialization entities.Specialization, err error) {
	err = row.Scan(&specialization.SpecializationID, &specialization.SpecializationName)
	return
}
