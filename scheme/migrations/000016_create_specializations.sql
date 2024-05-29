-- +goose Up
-- +goose StatementBegin

insert into tbl_specializations(specialization_name) values
                                                         ('Инженер'),
                                                         ('Инженер-программист'),
                                                         ('Бухгалтер'),
                                                         ('Техник-программист')

-- +goose StatementEnd