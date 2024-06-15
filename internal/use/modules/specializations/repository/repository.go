package repository

import (
	"context"
	"database/sql"
	"errors"

	"dogovorsBackEnd/internal/use/modules/specializations/entities"
)

var (
	ErrCannotCreate = errors.New("cannot create entity")
	ErrNotFound     = errors.New("not found")
	ErrCannotUpdate = errors.New("cannot update entity")
	ErrCannotDelete = errors.New("cannot delete entity")
)

type SpecializationsRepository interface {
	CreateSpecialization(ctx context.Context, specialization entities.Specialization) (id int, err error)
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
