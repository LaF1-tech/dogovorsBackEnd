package internal

import (
	"dogovorsBackEnd/internal/configs"
	"dogovorsBackEnd/internal/use"
	"dogovorsBackEnd/internal/use/types"
)

func Configure() (types.Runnable, types.Controllers, error) {
	db, err := configs.ConnectToDatabase(configs.Config.Database)
	if err != nil {
		return nil, types.Controllers{}, err
	}

	r, c := use.New(db)

	return r, c, nil
}
