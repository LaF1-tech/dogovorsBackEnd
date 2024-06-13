package entities

type ContractStatus string

var (
	ContractStatusAdded        ContractStatus = "Добавлено"
	ContractStatusChecked      ContractStatus = "Проверено"
	ContractStatusApproved     ContractStatus = "Утверждено"
	ContractStatusFulfilled    ContractStatus = "Исполнено"
	ContractStatusNotFulfilled ContractStatus = "Не исполнено"
)

type ContractExecutionControl struct {
	Id             int
	ContractId     int
	ContractStatus ContractStatus
}
