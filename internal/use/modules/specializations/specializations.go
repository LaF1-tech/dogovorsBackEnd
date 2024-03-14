package specializations

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/specializations/controller"
	"dogovorsBackEnd/internal/use/modules/specializations/controller/validator"
	"dogovorsBackEnd/internal/use/modules/specializations/handlers"
	"dogovorsBackEnd/internal/use/modules/specializations/repository"
)

func New(db *sql.DB) (handlers.SpecializationsHandler, controller.SpecializationsController) {
	repo := repository.New(db)

	v := validator.New()
	ctrl := controller.New(v, repo)

	h := handlers.New(ctrl)
	return h, ctrl
}
