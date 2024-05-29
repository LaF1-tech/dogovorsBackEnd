-- +goose Up
-- +goose StatementBegin

create table tbl_contracts
(
    contract_id     serial primary key,
    student_id      integer   NULL,
    employee_id     integer   NULL,
    template_id     integer   NULL,
    execution_date  timestamp NOT NULL,
    expiration_date timestamp NOT NULL,

    CONSTRAINT fk_student FOREIGN KEY (student_id) REFERENCES tbl_students (student_id) ON DELETE set null,
    CONSTRAINT fk_employee FOREIGN KEY (employee_id) REFERENCES tbl_employees (employee_id) ON DELETE set null,
    CONSTRAINT fk_template FOREIGN KEY (template_id) REFERENCES tbl_templates (template_id) ON DELETE set null

);

CALL register_updated_at_created_at_columns('tbl_contracts');

-- +goose StatementEnd