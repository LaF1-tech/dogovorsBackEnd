package dto

import (
	"time"

	"dogovorsBackEnd/internal/use/modules/contracts/entities"
)

type PatchContractRequestDTO struct {
	ContractID      int                      `json:"contract_id,omitempty"`
	StudentID       int                      `json:"student_id,omitempty"`
	EmployeeID      int                      `json:"employee_id,omitempty"`
	ApplicationType entities.ApplicationType `json:"application_type,omitempty"`
	ExecutionDate   time.Time                `json:"execution_date"`
	ExpirationDate  time.Time                `json:"expiration_date"`
}
