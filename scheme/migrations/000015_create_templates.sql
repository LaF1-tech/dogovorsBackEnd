-- +goose Up
-- +goose StatementBegin

insert into tbl_templates(template_name,
                          template_text,
                          necessary_data)
values ('Технологическая практика', '<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Договор на технологическую практику</title>
</head>
<body>
<div>
    <table border="1">
        <thead>
        <tr><td>Hello</td><td>World</td></tr>
        <tr><td>World</td><td>Hello</td></tr>
        </thead>
        <tbody>
        <tr><td>WOrld</td><td>Hello</td></tr>
        </tbody>
    </table>
</div>
</body>
</html>', 'a=>string,b=>integer,c=>date,d=>time,e=>datetime,f=>boolean'::hstore),
       ('Преддипломная практика', '<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Договор на преддипломную практику</title>
</head>
<body>
<div>
    <table border="1">
        <thead>
        <tr><td>Hello</td><td>World</td></tr>
        <tr><td>World</td><td>Hello</td></tr>
        </thead>
        <tbody>
        <tr><td>WOrld</td><td>Hello</td></tr>
        </tbody>
    </table>
</div>
</body>
</html>', 'a=>string,b=>integer'::hstore),
       ('Трудоустройство', '<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Договор на трудоустройство</title>
</head>
<body>
<div>
    <table border="1">
        <thead>
        <tr><td>Hello</td><td>World</td></tr>
        <tr><td>World</td><td>Hello</td></tr>
        </thead>
        <tbody>
        <tr><td>WOrld</td><td>Hello</td></tr>
        </tbody>
    </table>
</div>
</body>
</html>', 'b=>integer'::hstore)

-- +goose StatementEnd