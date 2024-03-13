-- +goose Up
-- +goose StatementBegin

create table tbl_employees
(
    employee_id serial primary key,
    username    VARCHAR(50) UNIQUE NOT NULL,
    password    TEXT               NOT NULL,
    first_name  VARCHAR(100)       NOT NULL,
    last_name   VARCHAR(100)       NOT NULL,
    permissions tp_permission[]    NULL
);

CALL register_updated_at_created_at_columns('tbl_employees');

-- +goose StatementEnd