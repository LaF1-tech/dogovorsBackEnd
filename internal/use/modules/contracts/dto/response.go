package dto

import (
	"dogovorsBackEnd/internal/use/modules/contracts/entities"
	"time"
)

type ContractItemResponseDTO struct {
	ContractID     int       `json:"contract_id,omitempty"`
	StudentID      int       `json:"student_id,omitempty"`
	EmployeeID     int       `json:"employee_id,omitempty"`
	TemplateID     int       `json:"template_id,omitempty"`
	ExecutionDate  time.Time `json:"execution_date"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type ContractAggregatedResponseDTO struct {
	ContractID        int                     `json:"contract_id,omitempty"`
	TemplateName      string                  `json:"template_name,omitempty"`
	StudentLastName   string                  `json:"student_last_name,omitempty"`
	StudentName       string                  `json:"student_name,omitempty"`
	StudentMiddleName string                  `json:"student_middle_name,omitempty"`
	ContractStatus    entities.ContractStatus `json:"contract_status,omitempty"`
	EmployeeFirstName string                  `json:"employee_first_name,omitempty"`
	EmployeeLastName  string                  `json:"employee_last_name,omitempty"`
	ExecutionDate     time.Time               `json:"execution_date"`
	ExpirationDate    time.Time               `json:"expiration_date"`
}

type ContractsViewResponseDTO struct {
	List []ContractAggregatedResponseDTO `json:"list,omitempty"`
}
