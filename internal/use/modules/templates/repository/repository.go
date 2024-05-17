package repository

import (
	"context"
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/templates/entities"
)

type TemplatesRepository interface {
	CreateTemplate(ctx context.Context, template entities.Template) (int, error)
	GetAllTemplatesPreview(ctx context.Context) ([]entities.TemplatePreview, error)
	GetTemplateByID(ctx context.Context, id int) (entities.Template, error)
	GetFullTemplateDataByContractID(ctx context.Context, contractID int) (entities.FullTemplate, error)
	PatchTemplateByID(ctx context.Context, template entities.Template) error
	DeleteTemplateByID(ctx context.Context, templateID int) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) TemplatesRepository {
	return &repository{db: db}
}
