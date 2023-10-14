CREATE TABLE IF NOT EXISTS atms
(
    id bigint NOT NULL PRIMARY KEY,
    address character varying(150),
    latitude real NOT NULL,
    longitude real NOT NULL,
    work_time_start time without time zone NOT NULL,
    work_time_end time without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS offices
(
    id bigint NOT NULL PRIMARY KEY,
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
    id bigint NOT NULL PRIMARY KEY,
    service_name character varying(250) NOT NULL
);


CREATE TABLE IF NOT EXISTS works_time_offices
(
    id bigint NOT NULL PRIMARY KEY,
    office_id bigint NOT NULL,
    weeks_day character varying(2) NOT NULL,
    work_time_start time without time zone NOT NULL,
    work_time_end time without time zone NOT NULL,
    lunch_time_start time without time zone,
    lunch_time_end time without time zone,
    is_legal boolean NOT NULL,
    FOREIGN KEY (office_id) 
        REFERENCES offices (id) 
        ON UPDATE NO ACTION 
        ON DELETE NO ACTION
);


CREATE TABLE IF NOT EXISTS offices_services
(
    id bigint NOT NULL PRIMARY KEY,
    office_id bigint NOT NULL,
    service_id bigint NOT NULL,
    FOREIGN KEY (office_id)
        REFERENCES offices (id)
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    FOREIGN KEY (service_id)
        REFERENCES services (id)
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);


CREATE TABLE IF NOT EXISTS charts_offices
(
    id bigint NOT NULL PRIMARY KEY,
    office_id bigint NOT NULL,
    weeks_day character varying(2) NOT NULL,
    charts_time time without time zone,
    chart integer NOT NULL,
    FOREIGN KEY (office_id)
        REFERENCES offices (id)
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS atms_services
(
    id bigint NOT NULL PRIMARY KEY,
    atm_id bigint NOT NULL,
    service_id bigint NOT NULL,
    FOREIGN KEY (atm_id)
        REFERENCES atms (id)
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    FOREIGN KEY (service_id)
        REFERENCES services (id)
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);