-- +goose Up
-- +goose StatementBegin

INSERT INTO tbl_users(username, password, first_name, last_name, permissions)
VALUES ('admin', '$2a$10$ETj.MjrlNhSyTMgQmnMrReDTSj019UX4vcrvFfMJ.sncy7ov596jG', 'admin', 'admin', '{Admin}')

-- +goose StatementEnd