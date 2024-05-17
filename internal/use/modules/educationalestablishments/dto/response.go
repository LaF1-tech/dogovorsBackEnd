package dto

type EducationalEstablishmentItemResponseDTO struct {
	EducationalEstablishmentID           int    `json:"educational_establishment_id,omitempty"`
	EducationalEstablishmentName         string `json:"educational_establishment_name,omitempty"`
	EducationalEstablishmentContactPhone string `json:"educational_establishment_contact_phone,omitempty"`
}

type EducationalEstablishmentsResponseDTO struct {
	List []EducationalEstablishmentItemResponseDTO `json:"list,omitempty"`
}
