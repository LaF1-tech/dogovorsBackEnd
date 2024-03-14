package educationalestablishments

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/educationalestablishments/controller"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/controller/validator"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/handlers"
	"dogovorsBackEnd/internal/use/modules/educationalestablishments/repository"
)

func New(db *sql.DB) (handlers.EducationalEstablishmentsHandler, controller.EducationalEstablishmentsController) {
	repo := repository.New(db)

	v := validator.New()
	ctrl := controller.New(v, repo)

	h := handlers.New(ctrl)
	return h, ctrl
}
