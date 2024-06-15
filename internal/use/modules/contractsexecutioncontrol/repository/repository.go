package repository

import (
	"context"
	"database/sql"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/entities"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

type ContractsExecutionControlRepository interface {
	GetContractsExecutionControlByContractID(ctx context.Context, id int) ([]entities.AggregatedContractExecutionControl, error)
	CreateContractExecutionControl(ctx context.Context, cec entities.ContractExecutionControl) (id int, err error)
	PatchContractExecutionControl(ctx context.Context, cec entities.ContractExecutionControl) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) ContractsExecutionControlRepository {
	return &repository{db: db}
}
