package entities

import "dogovorsBackEnd/internal/use/models"

type User struct {
	Timed

	UserID      int
	Username    string
	Password    string
	FirstName   string
	LastName    string
	Permissions []models.Permission
}
