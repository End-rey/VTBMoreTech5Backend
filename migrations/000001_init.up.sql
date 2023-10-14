CREATE TABLE IF NOT EXISTS atms
(
    id_atm bigint NOT NULL,
    address character varying(150) COLLATE pg_catalog."default",
    latitude real,
    longitude real,
    work_time_start time without time zone NOT NULL,
    work_time_end time without time zone NOT NULL,
    CONSTRAINT atms_pkey PRIMARY KEY (id_atm)
);

CREATE TABLE IF NOT EXISTS offices
(
    id_office bigint NOT NULL,
    salepointname character varying(150) COLLATE pg_catalog."default",
    address character varying(150) COLLATE pg_catalog."default",
    status boolean,
    officetype character varying(100) COLLATE pg_catalog."default",
    latitude real,
    longitude real,
    metrostation character varying(150) COLLATE pg_catalog."default",
    CONSTRAINT offices_pkey PRIMARY KEY (id_office)
);


CREATE TABLE IF NOT EXISTS services
(
    id_service bigint NOT NULL,
    servicename character varying(250) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT services_pkey PRIMARY KEY (id_service)
);


CREATE TABLE IF NOT EXISTS works_time_offices
(
    id bigint NOT NULL,
    id_office bigint NOT NULL,
    weeks_day character varying(2) COLLATE pg_catalog."default" NOT NULL,
    work_time_start time without time zone NOT NULL,
    work_time_end time without time zone NOT NULL,
    is_legal boolean NOT NULL,
    CONSTRAINT works_time_offices_pkey PRIMARY KEY (id),
    CONSTRAINT works_time_offices_id_office_fkey FOREIGN KEY (id_office)
        REFERENCES offices (id_office) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION

);


CREATE TABLE IF NOT EXISTS offices_services
(
    id bigint NOT NULL,
    office_id bigint NOT NULL,
    service_id bigint NOT NULL,
    CONSTRAINT offices_services_pk PRIMARY KEY (id),
    CONSTRAINT offices_services_office_id_fkey FOREIGN KEY (office_id)
        REFERENCES offices (id_office) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT offices_services_service_id_fkey FOREIGN KEY (service_id)
        REFERENCES services (id_service) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);


CREATE TABLE IF NOT EXISTS charts_offices
(
    id bigint NOT NULL,
    id_office bigint NOT NULL,
    weeks_day character varying(2) COLLATE pg_catalog."default" NOT NULL,
    charts_time time without time zone,
    chart integer NOT NULL,
    CONSTRAINT charts_offices_pk PRIMARY KEY (id),
    CONSTRAINT chart_offices_idoffice_fkey FOREIGN KEY (id_office)
        REFERENCES offices (id_office) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS atms_services
(
    id bigint NOT NULL,
    atm_id bigint NOT NULL,
    service_id bigint NOT NULL,
    CONSTRAINT atms_services_pk PRIMARY KEY (id),
    CONSTRAINT atms_services_atm_id_fkey FOREIGN KEY (atm_id)
        REFERENCES atms (id_atm) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT atms_services_service_id_fkey FOREIGN KEY (service_id)
        REFERENCES services (id_service) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);