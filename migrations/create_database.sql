create table tbl_users
(
    user_id    serial primary key,
    username   VARCHAR(50) UNIQUE NOT NULL,
    password   TEXT               NOT NULL,
    first_name VARCHAR(100),
    last_name  VARCHAR(100)
);

CREATE TYPE tp_application_status AS ENUM (
    'В процессе',
    'Обработано',
    'Отказано'
    );
CREATE TYPE tp_contract_relationship AS ENUM (
    'Прохождение произв практики',
    'Трудоустройство'
    );

CREATE TABLE tbl_applications
(
    application_id        SERIAL PRIMARY KEY,
    date_submitted        DATE                  NOT NULL,
    applicant_name        VARCHAR(255)          NOT NULL,
    applicant_surname     VARCHAR(255)          NOT NULL,
    applicant_middlename  VARCHAR(255)          NOT NULL,
    education_institution VARCHAR(255),
    specialization        VARCHAR(255),
    status                tp_application_status NOT NULL default ('В процессе')
);

CREATE TABLE tbl_contracts
(
    contract_id    SERIAL PRIMARY KEY,
    start_date     DATE                     NOT NULL,
    end_date       DATE,
    relationship   tp_contract_relationship NOT NULL default ('Прохождение произв практики'),
    application_id INTEGER REFERENCES tbl_applications (application_id)
);

create table tbl_archive_contracts
(
    json json
)