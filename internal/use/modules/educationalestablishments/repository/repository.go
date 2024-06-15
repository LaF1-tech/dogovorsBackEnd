package repository

import (
	"context"
	"database/sql"
	"errors"

	"dogovorsBackEnd/internal/use/modules/educationalestablishments/entities"
)

var (
	ErrCannotCreate = errors.New("cannot create entity")
	ErrNotFound     = errors.New("not found")
	ErrCannotUpdate = errors.New("cannot update entity")
	ErrCannotDelete = errors.New("cannot delete entity")
)

type EducationalEstablishmentsRepository interface {
	CreateEducationalEstablishment(ctx context.Context, establishment entities.EducationalEstablishment) (id int, err error)
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
