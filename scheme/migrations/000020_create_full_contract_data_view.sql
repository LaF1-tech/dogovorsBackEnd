-- +goose Up
-- +goose StatementBegin

create view vw_full_contract_data as
SELECT contract_id,
       ts.student_name,
       ts.student_last_name,
       te.first_name AS employee_first_name,
       te.last_name  AS employee_last_name,
       application_type,
       execution_date,
       expiration_date
FROM tbl_contracts
         INNER JOIN public.tbl_employees te on te.employee_id = tbl_contracts.employee_id
         inner join public.tbl_students ts on ts.student_id = tbl_contracts.student_id

-- +goose StatementEnd