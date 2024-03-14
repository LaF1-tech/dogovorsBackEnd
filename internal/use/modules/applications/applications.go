package applications

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/applications/controller"
	"dogovorsBackEnd/internal/use/modules/applications/controller/validator"
	"dogovorsBackEnd/internal/use/modules/applications/handlers"
	"dogovorsBackEnd/internal/use/modules/applications/repository"
)

func New(db *sql.DB) (handlers.ApplicationsHandler, controller.ApplicationsController) {
	repo := repository.New(db)

	v := validator.New()
	ctrl := controller.New(v, repo)

	h := handlers.New(ctrl)
	return h, ctrl
}
