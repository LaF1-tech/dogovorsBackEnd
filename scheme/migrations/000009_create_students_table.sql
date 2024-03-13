-- +goose Up
-- +goose StatementBegin

create table tbl_students
(
    student_id                   serial primary key,
    student_name                 varchar(255) NOT NULL,
    student_middle_name          varchar(255) NOT NULL,
    student_last_name            varchar(255) NOT NULL,
    specialization_id            integer      NULL,
    educational_establishment_id integer      NULL,

    CONSTRAINT fk_specializations FOREIGN KEY (specialization_id) REFERENCES tbl_specializations (specialization_id) ON DELETE set null,
    CONSTRAINT fk_educational_establishment FOREIGN KEY (educational_establishment_id) REFERENCES tbl_educational_establishments (educational_establishment_id) ON DELETE set null
);

CALL register_updated_at_created_at_columns('tbl_students');

-- +goose StatementEnd