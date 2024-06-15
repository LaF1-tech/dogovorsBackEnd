package repository

import (
	"context"
	"database/sql"
	"errors"

	"dogovorsBackEnd/internal/use/modules/auth/entities"
)

var (
	ErrCannotUpdate = errors.New("cannot update")
	ErrCannotDelete = errors.New("cannot delete")
	ErrCannotCreate = errors.New("cannot create")
)

type AuthRepository interface {
	GetUserByToken(ctx context.Context, token string) (entities.User, error)
	GetUserByUsername(ctx context.Context, username string) (entities.User, error)
	CreateUser(ctx context.Context, user entities.User) (int, error)
	UpdateUser(ctx context.Context, user entities.User) error
	DeleteUserByID(ctx context.Context, userID int) error

	CreateSession(ctx context.Context, session entities.Session) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) AuthRepository {
	return &repository{db: db}
}
