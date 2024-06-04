package repository

import (
	"context"
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/charts/entities"
)

type ChartsRepository interface {
	GetPeriodChart(ctx context.Context) ([]entities.PeriodChart, error)
	GetEducationalEstablishmentChart(ctx context.Context) ([]entities.EducationalEstablishmentCountChart, error)
	GetSpecializationsChart(ctx context.Context) ([]entities.SpecializationNameCountChart, error)
	GetTemplatesChart(ctx context.Context) ([]entities.TemplateNameCountChart, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) ChartsRepository {
	return &repository{db: db}
}
