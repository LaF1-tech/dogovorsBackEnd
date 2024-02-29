-- +goose Up
-- +goose StatementBegin

create table tbl_archive_contracts
(
    json json
);

CALL register_updated_at_created_at_columns('tbl_archive_contracts');

CREATE INDEX IF NOT EXISTS tbl_archive_contracts_created_at_idx ON tbl_archive_contracts USING hash (created_at);

-- +goose StatementEnd