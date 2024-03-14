package repository

import (
	"context"
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/specializations/entities"
)

type SpecializationsRepository interface {
	GetEducationalEstablishments(ctx context.Context) ([]entities.Specialization, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) SpecializationsRepository {
	return &repository{db: db}
}
