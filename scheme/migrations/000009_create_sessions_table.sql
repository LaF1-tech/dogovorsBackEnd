-- +goose Up
-- +goose StatementBegin

create table tbl_sessions
(
    user_id INTEGER REFERENCES tbl_users (user_id),
    token text UNIQUE NOT NULL primary key
);

CALL register_updated_at_created_at_columns('tbl_sessions');

-- +goose StatementEnd