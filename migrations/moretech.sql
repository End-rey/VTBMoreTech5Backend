-- Database: moretech

-- DROP DATABASE IF EXISTS moretech;

CREATE DATABASE moretech
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

-- Table: public.atm

-- DROP TABLE IF EXISTS public.atm;

-- Table: public.atm

-- DROP TABLE IF EXISTS public.atm;

CREATE TABLE IF NOT EXISTS public.atm
(
    id_atm bigint NOT NULL DEFAULT nextval('atm_id_atm_seq'::regclass),
    address character varying(150) COLLATE pg_catalog."default",
    latitude real,
    longitude real,
    work_time_start time without time zone NOT NULL,
    work_time_end time without time zone NOT NULL,
    CONSTRAINT atm_pkey PRIMARY KEY (id_atm)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.atm
    OWNER to postgres;
-- Table: public.office

-- DROP TABLE IF EXISTS public.office;

CREATE TABLE IF NOT EXISTS public.office
(
    id_office bigint NOT NULL DEFAULT nextval('office_id_office_seq'::regclass),
    salepointname character varying(150) COLLATE pg_catalog."default",
    address character varying(150) COLLATE pg_catalog."default",
    status boolean,
    officetype character varying(100) COLLATE pg_catalog."default",
    latitude real,
    longitude real,
    metrostation character varying(150) COLLATE pg_catalog."default",
    CONSTRAINT office_pkey PRIMARY KEY (id_office)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.office
    OWNER to postgres;


-- Table: public.service

-- DROP TABLE IF EXISTS public.service;

CREATE TABLE IF NOT EXISTS public.service
(
    id_service bigint NOT NULL DEFAULT nextval('service_idservice_seq'::regclass),
    servicename character varying(250) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT service_pkey PRIMARY KEY (id_service)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.service
    OWNER to postgres;


-- Table: public.timework_office

-- DROP TABLE IF EXISTS public.timework_office;

CREATE TABLE IF NOT EXISTS public.timework_office
(
    id_office bigint NOT NULL,
    days character varying(2) COLLATE pg_catalog."default" NOT NULL,
    timeworks character varying(15) COLLATE pg_catalog."default",
    CONSTRAINT timework_office_idoffice_fkey FOREIGN KEY (id_office)
        REFERENCES public.office (id_office) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.timework_office
    OWNER to postgres;

-- Table: public.office_service

-- DROP TABLE IF EXISTS public.office_service;

CREATE TABLE IF NOT EXISTS public.office_service
(
    office_id bigint NOT NULL,
    service_id bigint NOT NULL,
    CONSTRAINT office_service_pk PRIMARY KEY (office_id, service_id),
    CONSTRAINT office_service_office_id_fkey FOREIGN KEY (office_id)
        REFERENCES public.office (id_office) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT office_service_service_id_fkey FOREIGN KEY (service_id)
        REFERENCES public.service (id_service) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.office_service
    OWNER to postgres;

-- Table: public.chart_office

-- DROP TABLE IF EXISTS public.chart_office;

CREATE TABLE IF NOT EXISTS public.chart_office
(
    id_office bigint NOT NULL,
    dates character varying(2) COLLATE pg_catalog."default" NOT NULL,
    timechart time without time zone,
    chart integer NOT NULL,
    CONSTRAINT chart_office_idoffice_fkey FOREIGN KEY (id_office)
        REFERENCES public.office (id_office) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.chart_office
    OWNER to postgres;

-- Table: public.atm_service

-- DROP TABLE IF EXISTS public.atm_service;

CREATE TABLE IF NOT EXISTS public.atm_service
(
    atm_id bigint NOT NULL,
    service_id bigint NOT NULL,
    CONSTRAINT atm_service_pk PRIMARY KEY (atm_id, service_id),
    CONSTRAINT atm_service_atm_id_fkey FOREIGN KEY (atm_id)
        REFERENCES public.atm (id_atm) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT atm_service_service_id_fkey FOREIGN KEY (service_id)
        REFERENCES public.service (id_service) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.atm_service
    OWNER to postgres;