--liquibase formatted sql
--changeset Andrey Butusov:1
CREATE TABLE IF NOT EXISTS atms
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    creation_time timestamp(6),
    last_update_time timestamp(6),
    address character varying(150),
    latitude real NOT NULL,
    longitude real NOT NULL,
    work_time_start time without time zone NOT NULL,
    work_time_end time without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS offices
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    creation_time timestamp(6),
    last_update_time timestamp(6),
    sale_point_name character varying(150),
    address character varying(150),
    status boolean,
    office_type character varying(100),
    latitude real NOT NULL,
    longitude real NOT NULL,
    metro_station character varying(150)
);


CREATE TABLE IF NOT EXISTS services
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    creation_time timestamp(6),
    last_update_time timestamp(6),
    service_name character varying(250) NOT NULL
);


CREATE TABLE IF NOT EXISTS works_time_offices
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    creation_time timestamp(6),
    last_update_time timestamp(6),
    office_id bigint NOT NULL,
    weeks_day character varying(2),
    work_time_start time without time zone,
    work_time_end time without time zone,
    lunch_time_start time without time zone,
    lunch_time_end time without time zone,
    is_legal boolean,
    FOREIGN KEY (office_id) 
        REFERENCES offices  
);


CREATE TABLE IF NOT EXISTS offices_services
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    creation_time timestamp(6),
    last_update_time timestamp(6),
    office_id bigint NOT NULL,
    service_id bigint NOT NULL,
    FOREIGN KEY (office_id)
        REFERENCES offices,
    FOREIGN KEY (service_id)
        REFERENCES services
);


CREATE TABLE IF NOT EXISTS charts_offices
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    creation_time timestamp(6),
    last_update_time timestamp(6),
    office_id bigint NOT NULL,
    weeks_day character varying(2) NOT NULL,
    charts_time time without time zone,
    chart integer NOT NULL,
    FOREIGN KEY (office_id)
        REFERENCES offices
);

CREATE TABLE IF NOT EXISTS atms_services
(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    creation_time timestamp(6),
    last_update_time timestamp(6),
    atm_id bigint NOT NULL,
    service_id bigint NOT NULL,
    FOREIGN KEY (atm_id)
        REFERENCES atms,
    FOREIGN KEY (service_id)
        REFERENCES services
);

--rollback DROP TABLE IF EXISTS works_time_offices;

--rollback DROP TABLE IF EXISTS offices_services;

--rollback DROP TABLE IF EXISTS charts_offices;

--rollback DROP TABLE IF EXISTS atms_services;

--rollback DROP TABLE IF EXISTS atms;

--rollback DROP TABLE IF EXISTS offices;

--rollback DROP TABLE IF EXISTS services;