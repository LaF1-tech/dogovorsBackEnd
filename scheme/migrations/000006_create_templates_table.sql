-- +goose Up
-- +goose StatementBegin

create extension hstore;

create table tbl_templates
(
    template_id    serial primary key,
    template_name  varchar(255) NOT NULL,
    template_text  text         NOT NULL,
    necessary_data hstore NULL

);

CALL register_updated_at_created_at_columns('tbl_templates');

-- +goose StatementEnd