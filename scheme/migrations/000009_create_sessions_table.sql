-- +goose Up
-- +goose StatementBegin

create table tbl_sessions
(
    employee_id INTEGER REFERENCES tbl_employees (employee_id) ON DELETE CASCADE,
    token text UNIQUE NOT NULL primary key
);

CALL register_updated_at_created_at_columns('tbl_sessions');

-- +goose StatementEnd