CREATE DATABASE gochat_db
    WITH
    ENCODING = 'UTF-8';

CREATE ROLE admin WITH
    LOGIN
    SUPERUSER
    PASSWORD 'admin';

CREATE ROLE gc_api_service WITH
    LOGIN
    NOSUPERUSER
    INHERIT
    NOCREATEDB
    NOCREATEROLE
    NOREPLICATION
    PASSWORD 'gc-api-service';

ALTER ROLE gc_api_service SET search_path TO gochat_db;

\connect gochat_db

CREATE SCHEMA gochat_db
AUTHORIZATION gc_api_service;
