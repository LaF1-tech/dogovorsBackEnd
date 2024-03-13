package templates

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/templates/controller"
	"dogovorsBackEnd/internal/use/modules/templates/controller/validator"
	"dogovorsBackEnd/internal/use/modules/templates/handlers"
	"dogovorsBackEnd/internal/use/modules/templates/repository"
)

func New(db *sql.DB) (handlers.TemplateHandler, controller.TemplatesController) {
	repo := repository.New(db)

	v := validator.New()
	ctrl := controller.New(v, repo)

	h := handlers.New(ctrl)
	return h, ctrl
}
