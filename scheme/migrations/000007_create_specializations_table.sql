-- +goose Up
-- +goose StatementBegin

create table tbl_specializations
(
    specialization_id   serial primary key,
    specialization_name varchar(255) NOT NULL UNIQUE
);

CALL register_updated_at_created_at_columns('tbl_specializations');

-- +goose StatementEnd