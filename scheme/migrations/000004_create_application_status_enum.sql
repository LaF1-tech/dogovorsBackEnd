-- +goose Up
-- +goose StatementBegin

CREATE TYPE tp_application_status AS ENUM (
    'В процессе',
    'Обработано',
    'Отказано'
    );

-- +goose StatementEnd