package entities

import (
	"dogovorsBackEnd/internal/use/entities"
)

type Student struct {
	entities.Timed

	StudentID                  int
	SpecializationID           int
	EducationalEstablishmentID int
	StudentName                string
	StudentMiddleName          string
	StudentLastName            string
	StudentPhoneNumber         string
}
