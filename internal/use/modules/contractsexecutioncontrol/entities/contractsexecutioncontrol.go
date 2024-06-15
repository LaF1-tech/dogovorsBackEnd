package entities

import (
	"dogovorsBackEnd/internal/use/entities"
	"time"
)

type ContractExecutionControl struct {
	Id             int
	ContractId     int
	ControlDate    time.Time
	ContractStatus string

	entities.Timed
}

type AggregatedContractExecutionControl struct {
	Id              int
	ContractId      int
	TemplateName    string
	StudentName     string
	StudentLastName string
	ControlDate     time.Time
	ContractStatus  string
}
