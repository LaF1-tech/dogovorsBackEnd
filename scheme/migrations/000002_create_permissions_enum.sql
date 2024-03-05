-- +goose Up
-- +goose StatementBegin

create type tp_permission as enum
(
    'Admin'
);

-- +goose StatementEnd