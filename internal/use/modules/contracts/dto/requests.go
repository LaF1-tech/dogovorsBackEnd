package dto

import (
	"dogovorsBackEnd/internal/use/modules/contracts/entities"
)

type PatchContractRequestDTO struct {
	ContractID     int                     `json:"contract_id,omitempty"`
	ContractStatus entities.ContractStatus `json:"contract_status,omitempty"`
}
