package contractsexecutioncontrol

import (
	"database/sql"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/controller"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/controller/validator"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/handlers"
	"dogovorsBackEnd/internal/use/modules/contractsexecutioncontrol/repository"
)

func New(db *sql.DB) (handlers.ContractsExecutionControlHandler, controller.ContractsExecutionControlController) {
	repo := repository.New(db)

	v := validator.New()
	ctrl := controller.New(v, repo)

	h := handlers.New(ctrl)
	return h, ctrl
}
