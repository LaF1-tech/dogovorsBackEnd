package dto

import (
	"time"
)

type PatchContractRequestDTO struct {
	ContractID     int       `json:"contract_id,omitempty"`
	StudentID      int       `json:"student_id,omitempty"`
	EmployeeID     int       `json:"employee_id,omitempty"`
	TemplateID     int       `json:"application_type,omitempty"`
	ExecutionDate  time.Time `json:"execution_date"`
	ExpirationDate time.Time `json:"expiration_date"`
}
