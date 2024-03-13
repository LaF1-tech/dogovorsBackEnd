package entities

import (
	"dogovorsBackEnd/internal/use/entities"
	"dogovorsBackEnd/internal/use/models"
)

type User struct {
	entities.Timed

	UserID      int
	Username    string
	Password    string
	FirstName   string
	LastName    string
	Permissions []models.Permission
}
