package repository

import (
	"database/sql"
	"github.com/lib/pq"

	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/auth/entities"
	"dogovorsBackEnd/internal/use/utils/uslices"
)

func (r *repository) scanUser(row *sql.Row) (user entities.User, err error) {
	var permissions []string
	err = row.Scan(
		&user.UserID,
		&user.Username,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		pq.Array(&permissions),
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	user.Permissions = uslices.MapFunc(permissions, func(item string) models.Permission {
		return models.Permission(item)
	})

	return
}
