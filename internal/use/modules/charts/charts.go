package charts

import (
	"database/sql"
	"dogovorsBackEnd/internal/use/modules/charts/controller"
	"dogovorsBackEnd/internal/use/modules/charts/controller/validator"
	"dogovorsBackEnd/internal/use/modules/charts/handlers"
	"dogovorsBackEnd/internal/use/modules/charts/repository"
)

func New(db *sql.DB) (handlers.ChartsHandler, controller.ChartsController) {
	repo := repository.New(db)

	v := validator.New()
	ctrl := controller.New(v, repo)

	h := handlers.New(ctrl)
	return h, ctrl
}
