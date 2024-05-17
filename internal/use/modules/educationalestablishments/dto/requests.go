package dto

type CreateEducationalEstablishmentRequestDTO struct {
	EducationalEstablishmentName         string `json:"educational_establishment_name,omitempty"`
	EducationalEstablishmentContactPhone string `json:"educational_establishment_contact_phone,omitempty"`
}

type PatchEducationalEstablishmentRequestDTO struct {
	EducationalEstablishmentID           int    `json:"educational_establishment_id,omitempty"`
	EducationalEstablishmentName         string `json:"educational_establishment_name,omitempty"`
	EducationalEstablishmentContactPhone string `json:"educational_establishment_contact_phone,omitempty"`
}
