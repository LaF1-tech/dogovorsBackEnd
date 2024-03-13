-- +goose Up
-- +goose StatementBegin

create table tbl_contracts
(
    contract_id      serial primary key,
    student_id       integer                                                NULL,
    employee_id      integer                                                NULL,
    application_type tp_application default ('Прохождение произв практики') NOT NULL,
    execution_date   timestamp                                              NOT NULL,
    expiration_date  timestamp                                              NOT NULL,

    CONSTRAINT fk_student FOREIGN KEY (student_id) REFERENCES tbl_students (student_id) ON DELETE set null,
    CONSTRAINT fk_employee FOREIGN KEY (employee_id) REFERENCES tbl_employees (employee_id) ON DELETE set null

);

CALL register_updated_at_created_at_columns('tbl_contracts');

-- +goose StatementEnd