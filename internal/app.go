package internal

import "dogovorsBackEnd/internal/configs"

func App() error {
	db, err := configs.ConnectToDatabase(configs.Config.Database)
	if err != nil {
		return err
	}

	_ = db
	return nil
}
