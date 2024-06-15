package repository

import (
	"context"

	"dogovorsBackEnd/internal/use/modules/contracts/entities"
)

func (r *repository) GetAllContracts(ctx context.Context) ([]entities.AggregatedContract, error) {
	query := `
SELECT contract_id, 
       student_name, 
       student_last_name,
       employee_first_name, 
       employee_last_name,
       template_name,
       execution_date,
       expiration_date
from vw_full_contract_data order by contract_id
`
	res, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, ErrNotFound
	}
	var contracts []entities.AggregatedContract
	for res.Next() {
		contract, err := r.scanAggregatedContract(res)
		if err != nil {
			return nil, ErrNotFound
		}
		contracts = append(contracts, contract)
	}
	return contracts, nil
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
