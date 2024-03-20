-- +goose Up
-- +goose StatementBegin

create view vw_full_templates_data as
select tc.contract_id,
       tbl_templates.template_name,
       tbl_templates.template_content,
       tbl_templates.template_styles,
       tbl_templates.necessary_data,
       jsonb_build_object(
               'application_type', tc.application_type,
               'execution_date', tc.execution_date,
               'expiration_date', tc.expiration_date,
               'student_last_name', ts.student_last_name,
               'student_name', ts.student_name,
               'student_middle_name', ts.student_middle_name,
               'educational_establishment_name', tee.educational_establishment_name,
               'educational_establishment_contact_phone', tee.educational_establishment_contact_phone,
               'specialization_name', t.specialization_name,
               'employee_first_name', te.first_name,
               'employee_last_name', te.last_name
       ) || tcdr.data::jsonb as data
from tbl_templates
         inner join public.tbl_contract_data_registry tcdr on tbl_templates.template_id = tcdr.template_id
         inner join public.tbl_contracts tc on tc.contract_id = tcdr.contract_id
         inner join public.tbl_students ts on ts.student_id = tc.student_id
         inner join public.tbl_specializations t on t.specialization_id = ts.specialization_id
         inner join public.tbl_educational_establishments tee on tee.educational_establishment_id = ts.educational_establishment_id
         inner join public.tbl_employees te on te.employee_id = tc.employee_id;

-- +goose StatementEnd