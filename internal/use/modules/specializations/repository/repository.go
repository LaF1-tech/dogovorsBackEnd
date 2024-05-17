package repository

import (
	"context"
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/specializations/entities"
)

type SpecializationsRepository interface {
	CreateSpecialization(ctx context.Context, specialization entities.Specialization) (int, error)
	GetSpecializations(ctx context.Context) ([]entities.Specialization, error)
	GetSpecializationByID(ctx context.Context, id int) (entities.Specialization, error)
	PatchSpecializationByID(ctx context.Context, specialization entities.Specialization) error
	DeleteSpecializationByID(ctx context.Context, specID int) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) SpecializationsRepository {
	return &repository{db: db}
}
