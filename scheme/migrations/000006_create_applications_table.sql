-- +goose Up
-- +goose StatementBegin

CREATE TABLE tbl_applications
(
    application_id        SERIAL PRIMARY KEY,
    date_submitted        DATE                  NOT NULL,
    applicant_name        VARCHAR(255)          NOT NULL,
    applicant_surname     VARCHAR(255)          NOT NULL,
    applicant_middlename  VARCHAR(255)          NOT NULL,
    education_institution VARCHAR(255),
    specialization        VARCHAR(255),
    assigned_employer_id  INTEGER REFERENCES tbl_users (user_id),
    status                tp_application_status NOT NULL default ('В процессе')
);

CALL register_updated_at_created_at_columns('tbl_applications');

CREATE INDEX IF NOT EXISTS tbl_applications_created_at_idx ON tbl_applications USING hash (created_at);

-- +goose StatementEnd