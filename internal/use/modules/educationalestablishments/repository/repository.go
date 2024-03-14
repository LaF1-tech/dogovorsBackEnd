package repository

import (
	"context"
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/educationalestablishments/entities"
)

type EducationalEstablishmentsRepository interface {
	GetEducationalEstablishments(ctx context.Context) ([]entities.EducationalEstablishments, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) EducationalEstablishmentsRepository {
	return &repository{db: db}
}
