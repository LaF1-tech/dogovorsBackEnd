package repository

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/educationalestablishments/entities"
)

func (r *repository) scanEducationalEstablishment(row *sql.Rows) (ee entities.EducationalEstablishments, err error) {
	err = row.Scan(&ee.EducationalEstablishmentID, &ee.EducationalEstablishmentName, &ee.EducationalEstablishmentContactPhone)
	return
}
