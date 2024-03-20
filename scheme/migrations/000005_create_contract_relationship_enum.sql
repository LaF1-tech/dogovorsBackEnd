-- +goose Up
-- +goose StatementBegin

CREATE TYPE tp_application AS ENUM(
    'Технологическая практика',
    'Преддипломная практика',
    'Трудоустройство'
    );

-- +goose StatementEnd