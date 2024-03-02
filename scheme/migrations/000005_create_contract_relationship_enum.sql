-- +goose Up
-- +goose StatementBegin

CREATE TYPE tp_contract_relationship AS ENUM (
    'Прохождение произв практики',
    'Трудоустройство'
    );

-- +goose StatementEnd