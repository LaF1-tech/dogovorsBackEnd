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
       expiration_date from vw_full_contract_data
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

func (r *repository) GetPeriodAndCountForChart(ctx context.Context) ([]entities.ContractChart, error) {
	query := `SELECT DATE_TRUNC('month', execution_date) AS period,
       				 COUNT(*)                            AS contract_count
FROM tbl_contracts
GROUP BY period
ORDER BY period;
`
	response, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var data []entities.ContractChart
	for response.Next() {
		contractChartData, err := r.scanChartContract(response)
		if err != nil {
			return nil, err
		}
		data = append(data, contractChartData)
	}
	return data, err

}

func (r *repository) PatchContractByID(ctx context.Context, contract entities.Contract) error {
	query := `UPDATE tbl_contracts
SET student_id=COALESCE(NULLIF($1, 0), student_id),
    employee_id=COALESCE(NULLIF($2, 0), employee_id),
    template_id=COALESCE(NULLIF($3,0), template_id),
    execution_date=COALESCE(NULLIF($4, date('0000-00-00'))::timestamp, execution_date),
    expiration_date=COALESCE(NULLIF($5, date('0000-00-00'))::timestamp, expiration_date)
WHERE contract_id = $6
`
	_, err := r.db.ExecContext(ctx, query,
		contract.StudentID,
		contract.EmployeeID,
		contract.TemplateId,
		contract.ExecutionDate,
		contract.ExpirationDate,
		contract.ContractID)
	if err != nil {
		return err
	}
	return nil
}
