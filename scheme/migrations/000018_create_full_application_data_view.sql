-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE VIEW vw_full_application_data AS
SELECT application_id,
       tee.educational_establishment_name,
       ts.specialization_name,
       last_name,
       name,
       middle_name,
       phone_number,
       (SELECT json_agg(t.template_name)
        FROM json_each_text(types) AS elem(key, value)
                 JOIN tbl_templates t ON t.template_id = elem.key::integer) AS types,
       application_status,
       execution_date,
       expiration_date
from tbl_applications
         inner join public.tbl_educational_establishments tee
                    on tee.educational_establishment_id = tbl_applications.educational_establishment_id
         inner join public.tbl_specializations ts on ts.specialization_id = tbl_applications.specialization_id

-- +goose StatementEnd