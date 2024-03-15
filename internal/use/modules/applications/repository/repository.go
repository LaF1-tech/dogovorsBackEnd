package repository

import (
	"context"
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/applications/entities"
)

type ApplicationsRepository interface {
	CreateApplication(ctx context.Context, application entities.Application) (int, error)
	GetAllApplications(ctx context.Context) ([]entities.Application, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) ApplicationsRepository {
	return &repository{db: db}
}
