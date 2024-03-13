package entities

import "dogovorsBackEnd/internal/use/entities"

type Session struct {
	entities.Timed

	UserID int
	Token  string
}
