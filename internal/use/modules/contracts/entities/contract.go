package entities

import (
	"time"

	"dogovorsBackEnd/internal/use/entities"
)

type ApplicationType string

type Contract struct {
	entities.Timed

	ContractID     int
	StudentID      int
	EmployeeID     int
	TemplateId     int
	ExecutionDate  time.Time
	ExpirationDate time.Time
}
type AggregatedContract struct {
	ContractID        int
	StudentName       string
	StudentLastName   string
	EmployeeFirstName string
	EmployeeLastName  string
	TemplateName      string
	ExecutionDate     time.Time
	ExpirationDate    time.Time
}
