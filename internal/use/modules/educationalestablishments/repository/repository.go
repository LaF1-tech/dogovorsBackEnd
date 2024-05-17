package repository

import (
	"context"
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/educationalestablishments/entities"
)

type EducationalEstablishmentsRepository interface {
	CreateEducationalEstablishment(ctx context.Context, establishment entities.EducationalEstablishment) (int, error)
	GetEducationalEstablishments(ctx context.Context) ([]entities.EducationalEstablishment, error)
	GetEducationalEstablishmentByID(ctx context.Context, id int) (entities.EducationalEstablishment, error)
	PatchEducationalEstablishmentByID(ctx context.Context, establishment entities.EducationalEstablishment) error
	DeleteEducationalEstablishmentByID(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) EducationalEstablishmentsRepository {
	return &repository{db: db}
}
