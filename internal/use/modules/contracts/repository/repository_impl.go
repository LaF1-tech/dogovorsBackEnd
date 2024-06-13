package repository

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/contracts/entities"
)

func (r *repository) GetAllContracts(ctx context.Context) ([]entities.AggregatedContract, error) {
	query := `
SELECT contract_id, 
       template_name, 
       student_last_name,
       student_name, 
       student_middle_name, 
       CONTRACT_STATUS, 
       employee_first_name, 
       employee_last_name,
       execution_date,
       expiration_date
from vw_full_contract_execution_control_data
`
	res, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var contracts []entities.AggregatedContract
	for res.Next() {
		contract, err := r.scanAggregatedContract(res)
		if err != nil {
			return nil, err
		}
		contracts = append(contracts, contract)
	}
	return contracts, err
}

func (r *repository) GetContractByID(ctx context.Context, id int) (entities.Contract, error) {
	query := `
SELECT contract_id,
       student_id, 
       employee_id, 
       template_id, 
       execution_date,
       expiration_date from tbl_contracts WHERE contract_id=$1
`
	response := r.db.QueryRowContext(ctx, query, id)
	return r.scanContract(response)
}

func (r *repository) PatchContractByID(ctx context.Context, contract entities.ContractExecutionControl) error {
	query := `UPDATE tbl_contract_execution_control
SET CONTRACT_STATUS=COALESCE(NULLIF($1, '')::tp_execution_control_status, CONTRACT_STATUS)
    WHERE CONTRACT_ID = $2
`
	_, err := r.db.ExecContext(ctx, query,
		contract.ContractStatus,
		contract.ContractID)
	if err != nil {
		return err
	}
	return nil
}
