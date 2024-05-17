package dto

type CreateSpecializationRequestDTO struct {
	SpecializationName string `json:"specialization_name,omitempty"`
}

type PatchSpecializationRequestDTO struct {
	SpecializationID   int    `json:"specialization_id,omitempty"`
	SpecializationName string `json:"specialization_name,omitempty"`
}
