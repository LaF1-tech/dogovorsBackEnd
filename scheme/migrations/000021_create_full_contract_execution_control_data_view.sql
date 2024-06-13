-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE VIEW vw_full_contract_execution_control_data AS
SELECT tc.contract_id,
       tt.template_name,
       ts.student_last_name,
       ts.student_name,
       ts.student_middle_name,
       CONTRACT_STATUS,
       te.first_name as employee_first_name,
       te.last_name  as employee_last_name,
       tc.execution_date,
       tc.expiration_date
from tbl_contract_execution_control
         inner join public.tbl_contracts tc on tc.contract_id = tbl_contract_execution_control.CONTRACT_ID
         inner join public.tbl_students ts on ts.student_id = tc.student_id
         inner join public.tbl_templates tt on tt.template_id = tc.template_id
         inner join public.tbl_employees te on tc.employee_id = te.employee_id;


-- +goose StatementEnd