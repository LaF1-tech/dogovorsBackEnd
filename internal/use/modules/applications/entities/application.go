package entities

import (
	"time"

	"dogovorsBackEnd/internal/use/entities"
)

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
	LastName                   string
	Name                       string
	MiddleName                 string
	PhoneNumber                string
	Types                      map[string]map[string]any
	ApplicationStatus          ApplicationStatus
	ExecutionDate              time.Time
	ExpirationDate             time.Time
}
type AggregatedApplication struct {
	ApplicationID                int
	EducationalEstablishmentName string
	SpecializationName           string
	LastName                     string
	Name                         string
	MiddleName                   string
	PhoneNumber                  string
	Types                        map[string]map[string]any
	ApplicationStatus            ApplicationStatus
	ExecutionDate                time.Time
	ExpirationDate               time.Time
}
