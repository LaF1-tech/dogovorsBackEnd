package dto

import "time"

type AggregatedContractExecutionControlResponseDTO struct {
	Id              int       `json:"id,omitempty"`
	ContractId      int       `json:"contract_id,omitempty"`
	TemplateName    string    `json:"template_name,omitempty"`
	StudentLastName string    `json:"student_last_name,omitempty"`
	StudentName     string    `json:"student_name,omitempty"`
	ContractStatus  string    `json:"contract_status,omitempty"`
	ControlDate     time.Time `json:"control_date,omitempty"`
}

type ContractExecutionControlResponseDTO struct {
	Id             int       `json:"id,omitempty"`
	ContractId     int       `json:"contract_id,omitempty"`
	ContractStatus string    `json:"contract_status,omitempty"`
	ControlDate    time.Time `json:"control_date,omitempty"`
}

type AggregatedContractsExecutionControlResponseDTO struct {
	List []AggregatedContractExecutionControlResponseDTO `json:"list,omitempty"`
}
