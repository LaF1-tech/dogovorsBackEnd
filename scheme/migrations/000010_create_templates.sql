-- +goose Up
-- +goose StatementBegin

insert into tbl_templates(template_name,
                          template_content,
                          template_styles,
                          necessary_data)
values ('Технологическая практика', '<div class="doc_inside">
    <p>Открытое акционерное общество</p>
    <p>"Планар"</p>
    <p>ПРИКАЗ</p>
    <p>{{ .execution_date }} N 2</p>
    <p>г. Минск</p>
    <p>О зачислении на технологическую</p>
    <p>практику</p>
    <p>ПРИКАЗЫВАЮ:</p>
    <br>
    <br>
    <p>1. {{ .student_last_name }} {{ .student_name }} {{ .student_middle_name }}, студента(-ку) УО {{ .educational_establishment_name }},
        зачислить для прохождения технологической практики на период с {{ .execution_date }} по {{ .expiration_date }} с
        выполнением
        отдельных работ, предусмотренных должностными обязанностями {{ .specialization_name }}.</p>
    <p>2. Назначить руководителем технологической практики от организации главного {{ .specialization_name }}а
        _____________________.</p>
    <p>3. Назначить главного {{ .specialization_name }}а _______________ ответственным(-ой) за соблюдение _________
        требований
        безопасности труда.</p>
    <p>Основание: договор об организации производственного обучения от 01.12.2010 N 11/2010.</p>
    <br>
    <p>Директор _______ ____________</p>

    <p>С приказом ознакомлен(-а): _________ _____________
    </p>
    <p>
        {{ .execution_date }}
    </p>
    <br>
    <p>Примечание. Согласно п. 9 Положения N 860 студенты вузов во время прохождения технологической практики
        выполняют работу, предусмотренную разделом "Должностные обязанности" по соответствующей должности Единого
        квалификационного справочника должностей служащих.</p>
</div>', '* {
        margin: 12px;
        border: 0;
        padding: 0;
        vertical-align: top;
        font-family: Roboto;
    }

    .doc_inside {
        text-align: left;
        font-size: 12pt;
        display: inline-block;
        line-height: 170%;
    }

    @font-face {
        font-family: Roboto;
        src: url("fonts/RobotoFlex-Regular.ttf");
    }', null),
       ('Преддипломная практика', '<div class="doc_inside">
    <p>Открытое акционерное общество</p>
    <p>"Планар"</p>
    <p>ПРИКАЗ</p>
    <p>{{ .execution_date }} N 2</p>
    <p>г. Минск</p>
    <p>О зачислении на преддипломную</p>
    <p>практику</p>
    <p>ПРИКАЗЫВАЮ:</p>
    <br>
    <br>
    <p>1. {{ .student_last_name }} {{ .student_name }} {{ .student_middle_name }}, студента(-ку) УО {{ .educational_establishment_name }},
        зачислить для прохождения преддипломной практики на период с {{ .execution_date }} по {{ .expiration_date }} с
        выполнением
        отдельных работ, предусмотренных должностными обязанностями {{ .specialization_name }}.</p>
    <p>2. Назначить руководителем преддипломной практики от организации главного {{ .specialization_name }}а
        _____________________.</p>
    <p>3. Назначить главного {{ .specialization_name }}а _______________ ответственным(-ой) за соблюдение _________
        требований
        безопасности труда.</p>
    <p>Основание: договор об организации производственного обучения от 01.12.2010 N 11/2010.</p>
    <br>
    <p>Директор _______ ____________</p>

    <p>С приказом ознакомлен(-а): _________ _____________
    </p>
    <p>
        {{ .execution_date }}
    </p>
    <br>
    <p>Примечание. Согласно п. 9 Положения N 860 студенты вузов во время прохождения технологической практики
        выполняют работу, предусмотренную разделом "Должностные обязанности" по соответствующей должности Единого
        квалификационного справочника должностей служащих.</p>
</div>', '* {
        margin: 12px;
        border: 0;
        padding: 0;
        vertical-align: top;
        font-family: Roboto;
    }

    .doc_inside {
        text-align: left;
        font-size: 12pt;
        display: inline-block;
        line-height: 170%;
    }

    @font-face {
        font-family: Roboto;
        src: url("fonts/RobotoFlex-Regular.ttf");
    }', null),
       ('Трудоустройство', '<div>
    <p>Утвержден постановлением Совмина от 22.06.2011 N 821</p>
    <p>Форма</p>
    <p>__ _____________ 20__ г. _____________</p>
</div>
<div>
    <br>
    <br>
    <span>Гражданин _________________________________________________,</span>
    <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(Фамилия,
        собственное имя, отчество) (дата рождения)</p>
    <span>документ, удостоверяющий личность _________________________</span>
    <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(Серия
        (при наличии), номер, дата</p>
    <span>_____________________________________________________________</span>
    <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;выдачи,
        наименование государственного органа, его выдавшего,</p>
    <span>_____________________________________________________________,</span>
    <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;идентификационный
        номер (при наличии)</p>
    <span>Проживающий по адресу: _________________________________________</span>
    <p>__________________________________________________________________,</p>
    <span>С одной стороны, заказчик _________________________________________</span>
    <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(Наименование
        организации)</p>
    <p>__________________________________________________________________</p>
    <span>в лице ____________________________________________________________,</span>
    <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        (должность, фамилия, собственное имя, отчество)</p>
    <span>действующего на основании __________________________________________,</span>
    <p>с другой стороны, учреждение образования ___________________________________,</p>
    <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(наименование)</p>
    <p>______________________________________________________________________________</p>
    <span>в лице _____________________________________________________________________,</span>
    <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        (должность, фамилия, собственное имя, отчество)</p>
    <p>
        действующего на основании Устава, с третьей стороны, руководствуясь законодательством Республики Беларусь,
        заключили настоящий договор о нижеследующем:
    </p>
    <div>
        <p>1. <span>Гражданин _______________________________________________________</span>
        <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(Фамилия,
            собственное имя, отчество)</p>
        <p>__________________________________________________________________________ обязуется:</p>
        <ul>
            <li>1.1. пройти полный курс обучения сроком ______ лет в учреждении образования (направлению специальности,
                специализации) или квалификации (профессии рабочего, должности служащего)
                ____________________________________________________
            </li>
            <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(код
                и наименование</p>
            <li>
                ______________________________________________________________________
            </li>
            <p style="font-size:8pt;">специальности (направления специальности, специализации) или квалификации</p>
            <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(профессии
                рабочего, должности служащего)</p>
            <li>
                и получить квалификацию (профессию рабочего, должность служащего) __________
            </li>
            <li>
                _______________________________________________________________________________
            </li>
            <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(наименование
                квалификации (профессии рабочего, должности служащего)</p>
            <li>
                в соответствии с утвержденными учебными планами и программами на условиях, устанавливаемых в настоящем
                договоре;
            </li>
            <br>
            <li>
                1.2. после окончания учреждения образования в течение ____ лет отработать у заказчика на условиях,
                изложенных в настоящем договоре;
            </li>
            <li>
                1.3. письменно уведомить заказчика и учреждение образования о намерении расторгнуть настоящий договор с
                указанием причин и представлением подтверждающих их документов;
            </li>
            <li>
                1.4. возместить средства, затраченные государством на его подготовку, в республиканский и (или) местный
                бюджеты в случае расторжения договора в период получения образования при отсутствии оснований,
                установленных Правительством;
            </li>
            <li>
                1.5. возместить средства, затраченные государством на его подготовку, в республиканский и (или) местный
                бюджеты в порядке, определяемом Правительством Республики Беларусь.
            </li>
        </ul>
        <p>2. Заказчик обязуется:</p>
        <ul>
            <li>2.1. предоставить гражданину работу после окончания обучения в соответствии с полученной специальностью
                (направлением специальности, специализацией) и присвоенной квалификацией в (на) ________________________
                (место работы, наименование структурного подразделения) в должности (профессии)
                __________________________________________________;
            </li>
            <li>2.2. предоставить гражданину жилое помещение в соответствии с законодательством;</li>
            <li>2.3. в случае невозможности трудоустройства в соответствии с настоящим договором не позднее чем за два
                месяца до направления гражданина на работу информировать гражданина и учреждение образования о причинах
                расторжения или намерении изменить условия настоящего договора;
            </li>
            <li>2.4. возместить средства, затраченные государством на подготовку, в республиканский и (или) местный
                бюджеты в случае необоснованного расторжения или невыполнения условий настоящего договора в порядке,
                определяемом Правительством Республики Беларусь;
            </li>
            <li>2.5. создать условия для прохождения гражданином производственной практики, практики, производственного
                обучения, проведения практических занятий (производственного обучения) в соответствии с требованиями,
                установленными учебными планами и программами.
            </li>
        </ul>
        <p>3. Учреждение образования обязуется:</p>
        <ul>
            <li>3.1. обеспечить подготовку гражданина по специальности (направлению специальности, специализации),
                квалификации (профессии рабочего, должности служащего)
                ________________________________________________________________ (код и наименование специальности
                (направления специальности, специализации), квалификации (профессии рабочего, должности служащего) в
                соответствии с требованиями, установленными учебными планами и программами;
            </li>
            <li>3.2. направить гражданина после окончания учебы на работу в соответствии с подпунктом 2.1 пункта 2
                настоящего договора и уведомить об этом заказчика;
            </li>
            <li>3.3. уведомить заказчика об отчислении гражданина из учреждения образования с указанием причин.</li>
        </ul>
        <p>4. Стоимость обучения по настоящему договору составляет
            ______________________________________________________ белорусских рублей.</p>
        <span>Изменение стоимости обучения осуществляется в установленном законодательством порядке.</span>
        <p>5. Условия настоящего договора могут быть изменены по соглашению сторон в соответствии с законодательством
            Республики Беларусь.</p>
        <p>6. Настоящий договор действует со дня его подписания руководителем учреждения образования и до окончания
            срока обязательной работы.</p>
        <p>7. Дополнительные условия:</p>
        <p>
            ___________________________________________________________________________

            ___________________________________________________________________________

            ___________________________________________________________________________
        </p>
        <br>
        <table border="0">
            <thead>
            <tr>
                <th>
                    Заказчик
                </th>
                <th>
                    Руководитель учреждения образования
                </th>
                <th>
                    Гражданин
                </th>
            </tr>
            </thead>
            <tbody>
            <tr>
                <td>
                    _______________
                </td>
                <td>
                    ______________________________________
                </td>
                <td>
                    _______________
                </td>
            </tr>
            <tr style="text-align: center">
                <td>
                    (Подпись)
                </td>
                <td>
                    (Подпись)
                </td>
                <td>
                    (Подпись)
                </td>
            </tr>
            <tr>
                <td>
                    М.П.
                </td>
                <td>
                    М.П.
                </td>
            </tr>
            </tbody>
        </table>
        <br>
        <p> С заключением настоящего договора несовершеннолетним гражданином _____</p>
        <p>____________________________________________________________________________</p>
        <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(Фамилия,
            собственное имя, отчество)</p>
        <p>согласен ________________________________________________________________</p>
        <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(фамилия,
            собственное имя, отчество, степень родства,</p>
        <p>___________________________________________________________________________</p>
        <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;данные
            документа, удостоверяющего личность (серия (при наличии),</p>
        <p>___________________________________________________________________________</p>
        <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;номер,
            дата выдачи, наименование государственного органа,</p>
        <p>__________________________________________________________________________.</p>
        <p style="font-size:8pt;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;его
            выдавшего, идентификационный номер (при наличии)</p>
        <p>___________________</p>
        <p> (подпись)</p>
    </div>
</div>', 'div {
        font-family: Roboto;
        margin: 12px;
        font-size: 12pt;
    }

    th, td {
        padding-left: 24px;
        padding-right: 24px;
    }

    ul {
        list-style-type: none;
        padding-left: 0;
    }

    ul li {
        margin-bottom: 5px;
    }

    @font-face {
        font-family: Roboto;
        src: url("fonts/RobotoFlex-Regular.ttf");
    }', null)

-- +goose StatementEnd