package entities

import (
	"time"

	"dogovorsBackEnd/internal/use/entities"
)

type ContractStatus string

var (
	ContractStatusAdded        ContractStatus = "Добавлено"
	ContractStatusChecked      ContractStatus = "Проверено"
	ContractStatusApproved     ContractStatus = "Утверждено"
	ContractStatusFulfilled    ContractStatus = "Исполнено"
	ContractStatusNotFulfilled ContractStatus = "Не исполнено"
)

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
	TemplateName      string
	StudentLastName   string
	StudentName       string
	StudentMiddleName string
	ContractStatus    ContractStatus
	EmployeeFirstName string
	EmployeeLastName  string

	ExecutionDate  time.Time
	ExpirationDate time.Time
}

type ContractExecutionControl struct {
	Id             int
	ContractID     int
	ContractStatus ContractStatus
}
