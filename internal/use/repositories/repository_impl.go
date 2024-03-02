package repositories

import (
	"context"
	"dogovorsBackEnd/internal/use/entities"
	"github.com/lib/pq"
)

func (r *repository) GetUserByToken(ctx context.Context, token string) (entities.User, error) {
	query := `
select tu.user_id, tu.username, tu.password, tu.first_name, tu.last_name, tu.permissions, tu.created_at, tu.updated_at from tbl_sessions 
	 inner join public.tbl_users tu on tu.user_id = tbl_sessions.user_id
	 where token=$1
`

	row := r.db.QueryRowContext(ctx, query, token)
	return r.scanUser(row)
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (entities.User, error) {
	query := `
select user_id, username, password, first_name, last_name, permissions, created_at, updated_at from tbl_users 
	where username=$1
`
	row := r.db.QueryRowContext(ctx, query, username)
	return r.scanUser(row)
}

func (r *repository) CreateUser(ctx context.Context, user entities.User) (int, error) {
	query := `
insert into tbl_users(username, password, first_name, last_name, permissions) VALUES ($1,$2,$3,$4,$5) RETURNING user_id
`
	row := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.FirstName, user.LastName, pq.Array(user.Permissions))
	return r.scanId(row)
}

func (r *repository) CreateSession(ctx context.Context, session entities.Session) error {
	query := `
insert into tbl_sessions (user_id, token) values ($1, $2)
`

	_, err := r.db.ExecContext(ctx, query, session.UserID, session.Token)
	return err
}
