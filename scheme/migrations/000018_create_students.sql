-- +goose Up
-- +goose StatementBegin

INSERT INTO tbl_students (student_name, student_middle_name, student_last_name, specialization_id,
                          educational_establishment_id)
VALUES ('John', 'Arthur', 'Doe', 1, 1),
       ('Jane', 'Elizabeth', 'Smith', 2, 2),
       ('Michael', 'William', 'Johnson', 3, 3),
       ('Emily', 'Rose', 'Brown', 4, 4),
       ('David', 'Charles', 'Miller', 1, 5),
       ('Olivia', 'Anne', 'Wilson', 2, 6),
       ('Alexander', 'James', 'Taylor', 3, 7),
       ('Isabella', 'Marie', 'Thomas', 4, 8),
       ('Daniel', 'Joseph', 'Martin', 1, 9),
       ('Sophia', 'Grace', 'Davis', 1, 10),
       ('Ethan', 'Michael', 'Anderson', 2, 11),
       ('Emma', 'Louise', 'Robinson', 3, 12),
       ('Jacob', 'Christopher', 'Thompson', 4, 13),
       ('Mia', 'Nicole', 'Clark', 1, 14),
       ('Noah', 'Alexander', 'White', 2, 15),
       ('Abigail', 'Emily', 'Harris', 3, 16),
       ('William', 'Henry', 'Moore', 4, 17),
       ('Ava', 'Elizabeth', 'Jackson', 1, 18),
       ('Joshua', 'Thomas', 'Walker', 2, 19),
       ('Charlotte', 'Olivia', 'Taylor', 3, 20);

-- +goose StatementEnd