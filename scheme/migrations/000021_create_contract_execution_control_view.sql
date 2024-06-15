-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE VIEW VW_FULL_CONTRACT_EXECUTION_CONTROL AS
SELECT TCEC.id,
       TCEC.contract_id,
       TT.template_name,
       TS.student_last_name,
       TS.student_name,
       TCEC.contract_status,
       TCEC.control_date

FROM tbl_contract_execution_control TCEC
         INNER JOIN public.tbl_contracts tc on tc.contract_id = TCEC.contract_id
         INNER JOIN public.tbl_students ts on tc.student_id = ts.student_id
         INNER JOIN public.tbl_templates tt on tt.template_id = tc.template_id

-- +goose StatementEnd