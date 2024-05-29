-- +goose Up
-- +goose StatementBegin

create table tbl_educational_establishments
(
    educational_establishment_id            serial primary key,
    educational_establishment_name          varchar(255) NOT NULL,
    educational_establishment_contact_phone varchar(50)  NOT NULL
);

CALL register_updated_at_created_at_columns('tbl_educational_establishments');

-- +goose StatementEnd