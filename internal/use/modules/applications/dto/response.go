package dto

import "dogovorsBackEnd/internal/use/modules/applications/entities"

type CreateApplicationResponseDTO struct {
	ApplicationID              int                        `json:"application_id,omitempty"`
	EducationalEstablishmentID int                        `json:"educational_establishment_id,omitempty"`
	SpecializationID           int                        `json:"specialization_id,omitempty"`
	TemplateID                 int                        `json:"template_id,omitempty"`
	LastName                   string                     `json:"last_name,omitempty"`
	Name                       string                     `json:"name,omitempty"`
	MiddleName                 string                     `json:"middle_name,omitempty"`
	PhoneNumber                string                     `json:"phone_number,omitempty"`
	TemplateData               map[string]any             `json:"template_data,omitempty"`
	ApplicationStatus          entities.ApplicationStatus `json:"application_status,omitempty"`
}
