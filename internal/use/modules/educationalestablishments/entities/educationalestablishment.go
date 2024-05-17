package entities

import "dogovorsBackEnd/internal/use/entities"

type EducationalEstablishment struct {
	entities.Timed

	EducationalEstablishmentID           int
	EducationalEstablishmentName         string
	EducationalEstablishmentContactPhone string
}
