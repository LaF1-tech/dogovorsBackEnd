-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION check_control_date()
    RETURNS TRIGGER AS
$$
BEGIN
    PERFORM 1
    FROM tbl_contracts
    WHERE contract_id = NEW.contract_id
      AND NEW.control_date BETWEEN execution_date AND expiration_date;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'Дата должна быть в пределах даты исполнения и даты истечения договора';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE tbl_contract_execution_control
(
    ID              SERIAL PRIMARY KEY,
    CONTRACT_ID     INTEGER      NULL,
    CONTROL_DATE    timestamp    NOT NULL,
    CONTRACT_STATUS VARCHAR(255) NOT NULL DEFAULT ('Добавлено'),
    CONSTRAINT FK_CONTRACT FOREIGN KEY (CONTRACT_ID) REFERENCES tbl_contracts (contract_id) ON DELETE SET NULL
);

CALL register_updated_at_created_at_columns('tbl_contract_execution_control');

CREATE TRIGGER trg_check_control_date
    BEFORE INSERT OR UPDATE
    ON tbl_contract_execution_control
    FOR EACH ROW
EXECUTE FUNCTION check_control_date();

-- +goose StatementEnd