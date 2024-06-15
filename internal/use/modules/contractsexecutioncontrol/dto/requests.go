package dto

import "time"

type PatchContractExecutionControlRequestDTO struct {
	Id            int       `json:"id,omitempty"`
	ControlStatus string    `json:"contract_status,omitempty"`
	ControlDate   time.Time `json:"control_date,omitempty"`
}

type CreateContractExecutionControlRequestDTO struct {
	ContractId    int       `json:"contract_id,omitempty"`
	ControlDate   time.Time `json:"control_date,omitempty"`
	ControlStatus string    `json:"contract_status,omitempty"`
}
