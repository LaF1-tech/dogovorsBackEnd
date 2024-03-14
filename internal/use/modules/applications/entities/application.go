package entities

import "dogovorsBackEnd/internal/use/entities"

type ApplicationStatus string

const (
	ApplicationStatusInProgress ApplicationStatus = "В процессе"
	ApplicationStatusAccepted   ApplicationStatus = "Обработано"
	ApplicationStatusRejected   ApplicationStatus = "Отказано"
)

type Application struct {
	entities.Timed

	ApplicationID              int
	EducationalEstablishmentID int
	SpecializationID           int
	TemplateID                 int
	LastName                   string
	Name                       string
	MiddleName                 string
	PhoneNumber                string
	TemplateData               map[string]any
	ApplicationStatus          ApplicationStatus
}
