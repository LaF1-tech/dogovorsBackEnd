package repository

import (
	"database/sql"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/entities"
)

func (r *repository) scanAggregatedContractExecutionControl(row *sql.Rows) (cec entities.AggregatedContractExecutionControl, err error) {
	err = row.Scan(
		&cec.Id,
		&cec.ContractId,
		&cec.TemplateName,
		&cec.StudentLastName,
		&cec.StudentName,
		&cec.ContractStatus,
		&cec.ControlDate,
	)
	if err != nil {
		return entities.AggregatedContractExecutionControl{}, err
	}
	return
}
