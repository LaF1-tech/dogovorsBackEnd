-- +goose Up
-- +goose StatementBegin

insert into tbl_templates(template_name,
                          template_content,
                          template_styles,
                          necessary_data)
values ('Технологическая практика', '<body>
<h1>Договор на технологическую практику</h1>
<div>
    Студента {{ .student_last_name }} {{ .student_name }} {{ .student_middle_name }} <br>
    Обучающемуся в {{ .educational_establishment_name }} <br>
    По специальности {{ .specialization_name }} <br>
    Зачислить на технологическую практику в ОАО "Планар" <br>
    Дата исполнения: {{ .execution_date }} <br>
    Дата окончания срока действия: {{ .expiration_date }} <br>
</div>
<hr>
<div>Одобрил: {{ .employee_last_name }} {{ .employee_first_name }}</div>
<hr>
<br>
<br>
<br>
<br>
<br>
<div>М.П.</div>
</body>', '@font-face {
    font-family: Roboto;
    src: url("fonts/RobotoFlex-Regular.ttf");
}

body {
    font-family: Roboto;
}', null),
       ('Преддипломная практика', '<body>
<h1>Договор на преддипломную практику</h1>
<div>
    Студента {{ .student_last_name }} {{ .student_name }} {{ .student_middle_name }} <br>
    Обучающемуся в {{ .educational_establishment_name }} <br>
    По специальности {{ .specialization_name }} <br>
    Зачислить на преддпиломную практику в ОАО "Планар" <br>
    <br>
    <hr>
    Дата исполнения: {{ .execution_date }} <br>
    Дата окончания срока действия: {{ .expiration_date }} <br>
</div>
<hr>
<div>Одобрил: {{ .employee_last_name }} {{ .employee_first_name }}</div>
<hr>
<br>
<br>
<br>
<br>
<br>
<div>М.П.</div>
</body>', '@font-face {
    font-family: Roboto;
    src: url("fonts/RobotoFlex-Regular.ttf");
}

body {
    font-family: Roboto;
}', null),
       ('Трудоустройство', '<body>
<h1>Договор на трудоустройство</h1>
<div>
    Студента {{ .student_last_name }} {{ .student_name }} {{ .student_middle_name }} <br>
    Обучающемуся в {{ .educational_establishment_name }} <br>
    По специальности {{ .specialization_name }} <br>
    Зачислить в сотрудники ОАО "Планар" после окончания обучения<br>
    <br>
    <hr>
    Дата исполнения: {{ .execution_date }} <br>
    Дата окончания срока действия: {{ .expiration_date }} <br>
</div>
<hr>
<div>Одобрил: {{ .employee_last_name }} {{ .employee_first_name }}</div>
<hr>
<br>
<br>
<br>
<br>
<br>
<div>М.П.</div>
</body>', '@font-face {
    font-family: Roboto;
    src: url("fonts/RobotoFlex-Regular.ttf");
}

body {
    font-family: Roboto;
}', null)

-- +goose StatementEnd