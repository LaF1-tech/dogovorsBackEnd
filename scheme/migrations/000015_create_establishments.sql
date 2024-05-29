-- +goose Up
-- +goose StatementBegin

insert into tbl_educational_establishments(educational_establishment_name, educational_establishment_contact_phone)
values ('КБиП', '+375291234567'),
       ('БГУиР', '+375441234567'),
       ('УО "Минский государственный политехнический колледж"', '+375336543210'),
       ('Гродненский государственный университет', '+375291112233'),
       ('Минский инженерно-экономический колледж', '+375447778899'),
       ('Брестский государственный технический университет', '+375296666666'),
       ('Могилевский государственный университет имени А.А. Кулешова', '+375333445566'),
       ('Гомельский государственный технический университет имени П.О. Сухого', '+375442223344'),
       ('Белорусская государственная академия искусств', '+375297778899'),
       ('Витебский государственный медицинский университет', '+375294455566'),
       ('Институт бизнеса и технологий "Знание"', '+375298889900'),
       ('Молодечненский государственный университет', '+375445556677'),
       ('Белорусско-Российский университет', '+375443334455'),
       ('Гродненский государственный медицинский университет', '+375291234567'),
       ('Брестский государственный университет имени А.С. Пушкина', '+375446667788'),
       ('Минский государственный университет культуры', '+375295554433'),
       ('Полоцкий государственный университет', '+375444443322'),
       ('Могилевский государственный институт радиоэлектроники', '+375293332211'),
       ('Гомельский государственный университет имени Ф.Скорины', '+375291112233'),
       ('Белорусская государственная академия музыки', '+375449998877');


-- +goose StatementEnd