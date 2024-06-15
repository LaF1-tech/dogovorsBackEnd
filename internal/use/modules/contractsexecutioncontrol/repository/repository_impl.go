package repository

import (
	"context"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/entities"
	"dogovorsBackEnd/internal/use/utils/scanners"
)

func (r *repository) GetContractsExecutionControlByContractID(ctx context.Context, id int) ([]entities.AggregatedContractExecutionControl, error) {
	query := `
SELECT id, contract_id, template_name, student_last_name, student_name, contract_status, control_date FROM vw_full_contract_execution_control WHERE contract_id = $1 ORDER BY id
`
	res, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, ErrNotFound
	}
	var cecs []entities.AggregatedContractExecutionControl
	for res.Next() {
		cec, err := r.scanAggregatedContractExecutionControl(res)
		if err != nil {
			return nil, ErrNotFound
		}
		cecs = append(cecs, cec)
	}
	return cecs, nil
}

func (r *repository) CreateContractExecutionControl(ctx context.Context, cec entities.ContractExecutionControl) (id int, err error) {
	query := `INSERT INTO tbl_contract_execution_control (contract_id, control_date, contract_status)
VALUES ($1, $2, $3)
RETURNING id
`
	row := r.db.QueryRowContext(ctx, query, cec.ContractId, cec.ControlDate, cec.ContractStatus)
	return scanners.Id(row)
}

func (r *repository) PatchContractExecutionControl(ctx context.Context, cec entities.ContractExecutionControl) error {
	query := `UPDATE tbl_contract_execution_control
SET contract_status=COALESCE(NULLIF($1, ''), contract_status),
    control_date=COALESCE(NULLIF($2, '0001-01-01T00:00:00Z'::timestamp), control_date)
WHERE id = $3
`
	_, err := r.db.ExecContext(ctx, query, cec.ContractStatus, cec.ControlDate, cec.Id)
	if err != nil {
		return err
	}
	return nil
}
