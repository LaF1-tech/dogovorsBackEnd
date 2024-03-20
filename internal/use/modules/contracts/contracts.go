package contracts

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/contracts/controller"
	"dogovorsBackEnd/internal/use/modules/contracts/controller/validator"
	"dogovorsBackEnd/internal/use/modules/contracts/handlers"
	"dogovorsBackEnd/internal/use/modules/contracts/repository"
)

func New(db *sql.DB) (handlers.ContractsHandler, controller.ContractsController) {
	repo := repository.New(db)

	v := validator.New()
	ctrl := controller.New(v, repo)

	h := handlers.New(ctrl)
	return h, ctrl
}
