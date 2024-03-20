package entities

import "dogovorsBackEnd/internal/use/entities"

type ContractDataRegistry struct {
	entities.Timed

	TemplateID int
	ContractID int
	Data       map[string]any
}
