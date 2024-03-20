package dto

import (
	"time"

	"dogovorsBackEnd/internal/use/modules/contracts/entities"
)

type ContractItemResponseDTO struct {
	ContractID      int                      `json:"contract_id,omitempty"`
	StudentID       int                      `json:"student_id,omitempty"`
	EmployeeID      int                      `json:"employee_id,omitempty"`
	ApplicationType entities.ApplicationType `json:"application_type,omitempty"`
	ExecutionDate   time.Time                `json:"execution_date"`
	ExpirationDate  time.Time                `json:"expiration_date"`
}

type ContractAggregatedResponseDTO struct {
	ContractID        int                      `json:"contract_id,omitempty"`
	StudentName       string                   `json:"student_name,omitempty"`
	StudentLastName   string                   `json:"student_last_name,omitempty"`
	EmployeeFirstName string                   `json:"employee_first_name,omitempty"`
	EmployeeLastName  string                   `json:"employee_last_name,omitempty"`
	ApplicationType   entities.ApplicationType `json:"application_type,omitempty"`
	ExecutionDate     time.Time                `json:"execution_date"`
	ExpirationDate    time.Time                `json:"expiration_date"`
}

type ContractsViewResponseDTO struct {
	List []ContractAggregatedResponseDTO `json:"list,omitempty"`
}
