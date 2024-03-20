package repository

import (
	"context"
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/contracts/entities"
	tEntities "dogovorsBackEnd/internal/use/modules/templates/entities"
)

type ContractsRepository interface {
	GetAllContracts(ctx context.Context) ([]entities.AggregatedContract, error)
	GetContractByID(ctx context.Context, id int) (entities.Contract, error)
	PatchContractByID(ctx context.Context, contract entities.Contract) error

	GetAggregatedTemplateByContractID(ctx context.Context, id int) (tEntities.FullTemplate, error)
	GenerateTemplateContent(ctx context.Context, template tEntities.FullTemplate) (string, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) ContractsRepository {
	return &repository{db: db}
}
