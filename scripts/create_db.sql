CREATE DATABASE gochat_db
    WITH
    ENCODING = 'UTF-8';

CREATE ROLE admin WITH
    LOGIN
    SUPERUSER
    PASSWORD 'admin';

CREATE ROLE api_service_dev WITH
    LOGIN
    NOSUPERUSER
    INHERIT
    NOCREATEDB
    NOCREATEROLE
    NOREPLICATION
    PASSWORD 'gc-api-service';

ALTER ROLE api_service_dev SET search_path TO gochat_db;

\connect gochat_db

CREATE SCHEMA gochat_db
AUTHORIZATION api_service_dev;