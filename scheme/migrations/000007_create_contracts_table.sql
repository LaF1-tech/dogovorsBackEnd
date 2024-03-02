-- +goose Up
-- +goose StatementBegin

CREATE TABLE tbl_contracts
(
    contract_id    SERIAL PRIMARY KEY,
    start_date     DATE                     NOT NULL,
    end_date       DATE,
    relationship   tp_contract_relationship NOT NULL default ('Прохождение произв практики'),
    application_id INTEGER REFERENCES tbl_applications (application_id)
);

CALL register_updated_at_created_at_columns('tbl_contracts');

CREATE INDEX IF NOT EXISTS tbl_contracts_created_at_idx ON tbl_contracts USING hash (created_at);

-- +goose StatementEnd