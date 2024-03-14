package entities

import "dogovorsBackEnd/internal/use/entities"

type Specialization struct {
	entities.Timed

	SpecializationID   int
	SpecializationName string
}
