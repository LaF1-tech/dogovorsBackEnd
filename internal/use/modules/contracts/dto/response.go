package dto

import (
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
	ContractID        int       `json:"contract_id,omitempty"`
	StudentName       string    `json:"student_name,omitempty"`
	StudentLastName   string    `json:"student_last_name,omitempty"`
	EmployeeFirstName string    `json:"employee_first_name,omitempty"`
	EmployeeLastName  string    `json:"employee_last_name,omitempty"`
	TemplateName      string    `json:"template_name,omitempty"`
	ExecutionDate     time.Time `json:"execution_date"`
	ExpirationDate    time.Time `json:"expiration_date"`
}

type ContractChartResponseDTO struct {
	Period        time.Time `json:"period,omitempty"`
	ContractCount int       `json:"contract_count,omitempty"`
}

type ContractsChartResponseDTO struct {
	List []ContractChartResponseDTO `json:"list,omitempty"`
}

type ContractsViewResponseDTO struct {
	List []ContractAggregatedResponseDTO `json:"list,omitempty"`
}
