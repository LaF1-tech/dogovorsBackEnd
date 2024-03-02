-- +goose Up
-- +goose StatementBegin

create type tp_permission as enum
(
    'CreateUsers'
);

-- +goose StatementEnd