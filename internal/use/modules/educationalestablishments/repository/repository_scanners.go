package repository

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/educationalestablishments/entities"
)

func (r *repository) scanEducationalEstablishments(row *sql.Rows) (ee entities.EducationalEstablishment, err error) {
	err = row.Scan(&ee.EducationalEstablishmentID, &ee.EducationalEstablishmentName, &ee.EducationalEstablishmentContactPhone)
	return
}

func (r *repository) scanEducationalEstablishment(row *sql.Row) (ee entities.EducationalEstablishment, err error) {
	err = row.Scan(&ee.EducationalEstablishmentID, &ee.EducationalEstablishmentName, &ee.EducationalEstablishmentContactPhone)
	return
}
