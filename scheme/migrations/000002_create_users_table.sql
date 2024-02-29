-- +goose Up
-- +goose StatementBegin

create table tbl_users
(
    user_id    serial primary key,
    username   VARCHAR(50) UNIQUE NOT NULL,
    password   TEXT               NOT NULL,
    first_name VARCHAR(100),
    last_name  VARCHAR(100)
);

CALL register_updated_at_created_at_columns('tbl_users');

-- +goose StatementEnd