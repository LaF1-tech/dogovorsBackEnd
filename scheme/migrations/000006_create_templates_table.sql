-- +goose Up
-- +goose StatementBegin

create extension hstore;

create table tbl_templates
(
    template_id      serial primary key,
    template_name    varchar(255) NOT NULL,
    template_content text         NOT NULL,
    template_styles   text         NULL,
    necessary_data   hstore       NULL

);

CALL register_updated_at_created_at_columns('tbl_templates');

-- +goose StatementEnd