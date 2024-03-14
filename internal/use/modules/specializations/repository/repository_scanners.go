package repository

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/specializations/entities"
)

func (r *repository) scanSpecialization(row *sql.Rows) (ee entities.Specialization, err error) {
	err = row.Scan(&ee.SpecializationID, &ee.SpecializationName)
	return
}
