package repository

import (
	"context"
	"github.com/lib/pq"

	"dogovorsBackEnd/internal/use/modules/auth/entities"
	"dogovorsBackEnd/internal/use/utils/scanners"
)

func (r *repository) GetUserByUsername(ctx context.Context, username string) (entities.User, error) {
	query := `
SELECT employee_id, username, password, first_name, last_name, permissions, created_at, updated_at FROM tbl_employees
	WHERE username=$1
`
	row := r.db.QueryRowContext(ctx, query, username)
	return r.scanUser(row)
}

func (r *repository) CreateUser(ctx context.Context, user entities.User) (int, error) {
	query := `
INSERT INTO tbl_employees(username, password, first_name, last_name, permissions) VALUES ($1,$2,$3,$4,$5) RETURNING employee_id
`
	row := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.FirstName, user.LastName, pq.Array(user.Permissions))
	return scanners.Id(row)
}

func (r *repository) UpdateUser(ctx context.Context, user entities.User) error {
	query := `
UPDATE tbl_employees SET first_name=$2, last_name=$3, permissions=$4 WHERE employee_id = $1
`

	_, err := r.db.ExecContext(ctx, query, user.UserID, user.FirstName, user.LastName, pq.Array(user.Permissions))
	if err != nil {
		return err
	}

	if user.Password != "" {
		query = `
UPDATE tbl_employees SET password = $2 WHERE employee_id = $1
`

		_, err = r.db.ExecContext(ctx, query, user.UserID, user.Password)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) DeleteUserByID(ctx context.Context, userID int) error {
	query := `
DELETE FROM tbl_employees WHERE employee_id = $1
`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

func (r *repository) GetUserByToken(ctx context.Context, token string) (entities.User, error) {
	query := `
SELECT tu.employee_id, tu.username, tu.password, tu.first_name, tu.last_name, tu.permissions, tu.created_at, tu.updated_at FROM tbl_sessions 
	 INNER JOIN public.tbl_employees tu ON tu.employee_id = tbl_sessions.employee_id
	 WHERE token=$1
`

	row := r.db.QueryRowContext(ctx, query, token)
	return r.scanUser(row)
}

func (r *repository) CreateSession(ctx context.Context, session entities.Session) error {
	query := `
INSERT INTO tbl_sessions (employee_id, token) VALUES ($1, $2)
`

	_, err := r.db.ExecContext(ctx, query, session.UserID, session.Token)
	return err
}
