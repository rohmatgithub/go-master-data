-- +migrate Up

-- +migrate StatementBegin


CREATE TABLE IF NOT EXISTS country (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(20)     NOT NULL UNIQUE,
    name            VARCHAR(100)    NOT NULL,
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted         BOOLEAN
);

CREATE TABLE IF NOT EXISTS province (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(20)     NOT NULL UNIQUE,
    name            VARCHAR(100)    NOT NULL,
    parent_id       INTEGER,
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS district (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(20)     NOT NULL UNIQUE,
    name            VARCHAR(100)    NOT NULL,
    parent_id       INTEGER,
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted         BOOLEAN
);

CREATE TABLE IF NOT EXISTS sub_district (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(20)     NOT NULL UNIQUE,
    name            VARCHAR(100)    NOT NULL,
    parent_id       INTEGER,
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted         BOOLEAN
);

CREATE TABLE IF NOT EXISTS urban_village (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(20)     NOT NULL UNIQUE,
    name            VARCHAR(100)    NOT NULL,
    parent_id       INTEGER,
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted         BOOLEAN
);


-- +migrate StatementEnd