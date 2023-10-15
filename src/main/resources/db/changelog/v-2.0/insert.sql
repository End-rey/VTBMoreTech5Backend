INSERT INTO services (creation_time, last_update_time, service_name)
VALUES
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Кредит наличными'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Экспресс кредит'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Рефинансирование'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Карты'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Ипотека'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Автокредиты'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Вклады и счета'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Инвестиции'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Платежи и переводы'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Страховые и сервисные продукты'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Аренда сейфовых ячеек'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Банкротство физических лиц'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Валюта и золото'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Электронные банковские услуги'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Пенсионное обслуживание'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Управление активами'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Корпоративное финансирование'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Лизинг'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Доверительное управление'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Кредитные карты для бизнеса'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Брокерские услуги'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Онлайн-курсы финансового образования'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Услуги для иностранных клиентов'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Финансовое планирование и консультации'),
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Экологически устойчивые финансовые продукты')
;


INSERT INTO charts_offices (creation_time, last_update_time, office_id, weeks_day, charts_time, chart)
SELECT
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    o.id AS office_id,
    day AS weeks_day,
    make_time(hour, 0, 0) AS charts_time,
    floor(random() * 100)
FROM
    offices o,
    generate_series(0, 6) AS day,
    generate_series(0, 23) AS hour;
	