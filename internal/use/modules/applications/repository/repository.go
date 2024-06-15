package repository

import (
	"context"
	"database/sql"
	"errors"

	"dogovorsBackEnd/internal/use/modules/applications/entities"
	entities2 "dogovorsBackEnd/internal/use/modules/contracts/entities"
)

var (
	ErrNotFound     = errors.New("entity not found")
	ErrCannotCreate = errors.New("cannot create entity")
)

type ApplicationsRepository interface {
	CreateApplication(ctx context.Context, application entities.Application) (int, error)
	GetAllApplications(ctx context.Context) ([]entities.AggregatedApplication, error)
	GetApplicationByID(ctx context.Context, id int) (entities.Application, error)
	PatchApplicationByID(ctx context.Context, application entities.Application) error

	CreateContract(ctx context.Context, contract entities2.Contract) (int, error)
	CreateContractExecutionControl(ctx context.Context, cec entities.ContractExecutionControl) (int, error)
	CreateStudent(ctx context.Context, student entities.Student) (int, error)
	CreateContractDataRegistryEntry(ctx context.Context, cdr entities.ContractDataRegistry) (int, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) ApplicationsRepository {
	return &repository{db: db}
}
