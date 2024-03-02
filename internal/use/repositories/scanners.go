package repositories

import (
	"database/sql"
	"dogovorsBackEnd/internal/use/entities"
	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/utils/uslices"
	"github.com/lib/pq"
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

func (r *repository) scanId(row *sql.Row) (id int, err error) {
	err = row.Scan(&id)
	return
}
