package repository

import (
	"context"
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/templates/entities"
)

type TemplatesRepository interface {
	GetAllTemplatesPreview(ctx context.Context) ([]entities.TemplatePreview, error)
	GetTemplateByID(ctx context.Context, id int) (entities.Template, error)
	GetFullTemplateDataByContractID(ctx context.Context, contractID int) (entities.FullTemplate, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) TemplatesRepository {
	return &repository{db: db}
}
