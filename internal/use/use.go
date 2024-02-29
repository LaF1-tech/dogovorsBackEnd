package use

import (
	"database/sql"
	"dogovorsBackEnd/internal/use/controllers"
	"dogovorsBackEnd/internal/use/controllers/validators"
	"dogovorsBackEnd/internal/use/handlers"
	"dogovorsBackEnd/internal/use/repositories"
)

func New(database *sql.DB) (handlers.Handler, controllers.Controller, error) {
	repo := repositories.New(database)

	validator := validators.New()
	ctrl := controllers.New(validator, repo)

	h := handlers.New(ctrl)

	return h, ctrl, nil
}
