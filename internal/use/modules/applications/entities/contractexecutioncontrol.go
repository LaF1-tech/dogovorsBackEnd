package entities

import "time"

type ContractExecutionControl struct {
	Id             int
	ContractId     int
	ControlDate    time.Time
	ContractStatus string
}
