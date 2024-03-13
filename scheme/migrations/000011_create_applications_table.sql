-- +goose Up
-- +goose StatementBegin

create table tbl_applications
(
    application_id               serial primary key,
    educational_establishment_id integer                                      NULL,
    specialization_id            integer                                      NULL,
    template_id                  integer                                      NULL,
    last_name                    varchar(255)                                 NOT NULL,
    name                         varchar(255)                                 NOT NULL,
    middle_name                  varchar(255)                                 NOT NULL,
    phone_number                 varchar(50)                                  NOT NULL,
    template_data                json                                         NOT NULL,
    application_status           tp_application_status default ('В процессе') NOT NULL,

    CONSTRAINT fk_specializations FOREIGN KEY (specialization_id) REFERENCES tbl_specializations (specialization_id) ON DELETE set null,
    CONSTRAINT fk_educational_establishment FOREIGN KEY (educational_establishment_id) REFERENCES tbl_educational_establishments (educational_establishment_id) ON DELETE set null,
    CONSTRAINT fk_template FOREIGN KEY (template_id) REFERENCES tbl_templates (template_id) ON DELETE set null

);

CALL register_updated_at_created_at_columns('tbl_applications');

-- +goose StatementEnd