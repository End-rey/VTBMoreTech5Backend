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

-- Генерация данных для таблицы atms
INSERT INTO atms (creation_time, last_update_time, address, latitude, longitude, work_time_start, work_time_end)
SELECT
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    'Улица ' || floor(random() * 1000)::integer || ', Дом ' || floor(random() * 100)::integer,
    random() * (56.094033 - 56.039295) + 56.039295,
    random() * (38.135920 - 37.071069) + 37.071069,
    make_time(floor(random() * 24)::integer, floor(random() * 60)::integer, floor(random() * 60)::integer), -- Время начала работы
    make_time(floor(random() * 24)::integer, floor(random() * 60)::integer, floor(random() * 60)::integer)
FROM generate_series(1, 200) AS i;

-- Генерация данных для таблицы offices
INSERT INTO offices (creation_time, last_update_time, sale_point_name, address, status, office_type, latitude, longitude, metro_station)
SELECT
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    'Отделение банка ' || i,
    'Улица ' || floor(random() * 1000)::integer || ', Дом ' || floor(random() * 100)::integer,
    random() > 0.5,  -- Случайный статус true/false
    'Тип отделения ' || i,
    random() * (56.094033 - 56.039295) + 56.039295,
    random() * (38.135920 - 37.071069) + 37.071069,
    'Метро ' || i
FROM generate_series(1, 200) AS i;

INSERT INTO charts_offices (creation_time, last_update_time, office_id, weeks_day, charts_time, chart)
SELECT
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    o.id AS office_id,
    day AS weeks_day,
    make_time(hour, 0, 0) AS charts_time,
    floor(random() * 100)  -- Здесь можно установить диапазон загруженности
FROM
    offices o,
    generate_series(0, 6) AS day,  -- 0 - Понедельник, 1 - Вторник, и так далее
    generate_series(0, 23) AS hour;
	
INSERT INTO offices_services (creation_time, last_update_time, office_id, service_id)
SELECT
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    o.id AS office_id,
    floor(random() * 25) + 126 AS service_id
FROM
    offices o
LIMIT 200;

-- Генерация данных для таблицы atms_services (для банкоматов)
INSERT INTO atms_services (creation_time, last_update_time, atm_id, service_id)
SELECT
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    a.id AS atm_id,
    floor(random() * 25) + 126 AS service_id
FROM
    atms a
LIMIT 200;

INSERT INTO works_time_offices (creation_time, last_update_time, office_id, weeks_day, work_time_start, work_time_end, lunch_time_start, lunch_time_end, is_legal)
SELECT
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    o.id AS office_id,
    wd AS weeks_day,
    make_time(floor(random() * 24)::integer, floor(random() * 60)::integer, 0) AS work_time_start,
    make_time(floor(random() * 24)::integer, floor(random() * 60)::integer, 0) AS work_time_end,
    make_time(floor(random() * 24)::integer, floor(random() * 60)::integer, 0) AS lunch_time_start,
    make_time(floor(random() * 24)::integer, floor(random() * 60)::integer, 0) AS lunch_time_end,
    random() > 0.5  -- Случайный статус true/false
FROM
    offices o,
    (SELECT generate_series(0, 6) AS wd) AS week_days;