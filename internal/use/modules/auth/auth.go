package auth

import (
	"database/sql"

	"dogovorsBackEnd/internal/use/modules/auth/controller"
	"dogovorsBackEnd/internal/use/modules/auth/controller/validator"
	"dogovorsBackEnd/internal/use/modules/auth/handlers"
	"dogovorsBackEnd/internal/use/modules/auth/repository"
)

func New(db *sql.DB) (handlers.AuthHandler, controller.AuthController) {
	repo := repository.New(db)

	v := validator.New()
	ctrl := controller.New(v, repo)

	h := handlers.New(ctrl)
	return h, ctrl
}
