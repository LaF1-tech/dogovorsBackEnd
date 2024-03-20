package entities

import (
	"time"

	"dogovorsBackEnd/internal/use/entities"
)

type ApplicationType string

const (
	ApplicationTypeInternship            ApplicationType = "Технологическая практика"
	ApplicationTypeUndergraduatePractice ApplicationType = "Преддипломная практика"
	ApplicationTypeEmployment            ApplicationType = "Трудоустройство"
)

func ApplicationTypeById(id int) (t ApplicationType) {
	switch id {
	case 1:
		t = ApplicationTypeInternship
		break
	case 2:
		t = ApplicationTypeUndergraduatePractice
		break
	case 3:
		t = ApplicationTypeEmployment
		break
	}

	return
}

type Contract struct {
	entities.Timed

	ContractID      int
	StudentID       int
	EmployeeID      int
	ApplicationType ApplicationType
	ExecutionDate   time.Time
	ExpirationDate  time.Time
}
type AggregatedContract struct {
	ContractID        int
	StudentName       string
	StudentLastName   string
	EmployeeFirstName string
	EmployeeLastName  string
	ApplicationType   ApplicationType
	ExecutionDate     time.Time
	ExpirationDate    time.Time
}
