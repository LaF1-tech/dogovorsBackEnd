package repositories

import (
	"context"
	"database/sql"
	"dogovorsBackEnd/internal/use/entities"
)

type Repository interface {
	GetUserByToken(ctx context.Context, token string) (entities.User, error)
	GetUserByUsername(ctx context.Context, username string) (entities.User, error)
	CreateUser(ctx context.Context, user entities.User) (int, error)

	CreateSession(ctx context.Context, session entities.Session) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}
