package internal

import (
	"dogovorsBackEnd/internal/configs"
	"dogovorsBackEnd/internal/use"
	"dogovorsBackEnd/internal/use/controllers"
	"dogovorsBackEnd/internal/use/handlers"
)

func Configure() (handlers.Handler, controllers.Controller, error) {
	db, err := configs.ConnectToDatabase(configs.Config.Database)
	if err != nil {
		return nil, nil, err
	}

	return use.New(db)
}
