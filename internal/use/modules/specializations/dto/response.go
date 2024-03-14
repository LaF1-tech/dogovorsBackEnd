package dto

type SpecializationItemResponseDTO struct {
	SpecializationID   int    `json:"specialization_id,omitempty"`
	SpecializationName string `json:"specialization_name,omitempty"`
}

type SpecializationsResponseDTO struct {
	List []SpecializationItemResponseDTO `json:"list,omitempty"`
}
