package dto

import (
	"dogovorsBackEnd/internal/use/modules/applications/entities"
	"time"
)

type ApplicationItemResponseDTO struct {
	ApplicationID              int                        `json:"id,omitempty"`
	EducationalEstablishmentID int                        `json:"educational_establishment_id,omitempty"`
	SpecializationID           int                        `json:"specialization_id,omitempty"`
	LastName                   string                     `json:"last_name,omitempty"`
	Name                       string                     `json:"name,omitempty"`
	MiddleName                 string                     `json:"middle_name,omitempty"`
	PhoneNumber                string                     `json:"phone_number,omitempty"`
	Types                      map[int]map[string]any     `json:"types,omitempty"`
	ApplicationStatus          entities.ApplicationStatus `json:"application_status,omitempty"`
	ExecutionDate              time.Time                  `json:"execution_date,omitempty"`
	ExpirationDate             time.Time                  `json:"expiration_date,omitempty"`
}

type ApplicationViewResponseDTO struct {
	ApplicationID                int                        `json:"id,omitempty"`
	EducationalEstablishmentName string                     `json:"educational_establishment_name,omitempty"`
	SpecializationName           string                     `json:"specialization_name,omitempty"`
	LastName                     string                     `json:"last_name,omitempty"`
	Name                         string                     `json:"name,omitempty"`
	MiddleName                   string                     `json:"middle_name,omitempty"`
	PhoneNumber                  string                     `json:"phone_number,omitempty"`
	Types                        map[int]map[string]any     `json:"types,omitempty"`
	ApplicationStatus            entities.ApplicationStatus `json:"application_status,omitempty"`
	ExecutionDate                time.Time                  `json:"execution_date,omitempty"`
	ExpirationDate               time.Time                  `json:"expiration_date,omitempty"`
}

type ApplicationsViewResponseDTO struct {
	List []ApplicationViewResponseDTO `json:"list,omitempty"`
}
