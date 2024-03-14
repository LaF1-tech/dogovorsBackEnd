package entities

import "dogovorsBackEnd/internal/use/entities"

type EducationalEstablishments struct {
	entities.Timed

	EducationalEstablishmentID           int
	EducationalEstablishmentName         string
	EducationalEstablishmentContactPhone string
}
