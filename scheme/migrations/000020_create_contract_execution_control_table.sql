-- +goose Up
-- +goose StatementBegin

CREATE TABLE tbl_contract_execution_control
(
    ID              SERIAL PRIMARY KEY,
    CONTRACT_ID     INTEGER                     NULL,
    CONTRACT_STATUS tp_execution_control_status NOT NULL DEFAULT ('Добавлено'),
    CONSTRAINT FK_CONTRACT FOREIGN KEY (CONTRACT_ID) REFERENCES tbl_contracts (contract_id) ON DELETE SET NULL
);

CALL register_updated_at_created_at_columns('tbl_contract_execution_control');

-- +goose StatementEnd