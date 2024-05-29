-- +goose Up
-- +goose StatementBegin

create table tbl_contract_data_registry
(
    template_id integer NULL,
    contract_id integer NULL,
    data        json    NOT NULL,


    CONSTRAINT fk_template FOREIGN KEY (template_id) REFERENCES tbl_templates (template_id) ON DELETE CASCADE,
    CONSTRAINT fk_contract FOREIGN KEY (contract_id) REFERENCES tbl_contracts (contract_id) ON DELETE CASCADE

);

CALL register_updated_at_created_at_columns('tbl_contract_data_registry');

-- +goose StatementEnd