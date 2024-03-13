-- +goose Up
-- +goose StatementBegin

CREATE TYPE tp_application AS ENUM(
    'Прохождение произв практики',
    'Прохождение пред дипл практики',
    'Трудоустройство'
    );

-- +goose StatementEnd