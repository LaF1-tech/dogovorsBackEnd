-- +goose Up
-- +goose StatementBegin

create type tp_execution_control_status as enum
    (
        'Добавлено',
        'Проверено',
        'Утверждено',
        'Исполнено',
        'Неисполнено'
        );

-- +goose StatementEnd