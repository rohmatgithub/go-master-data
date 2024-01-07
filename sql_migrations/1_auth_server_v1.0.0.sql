-- +migrate Up

-- +migrate StatementBegin


CREATE TABLE IF NOT EXISTS country (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(20)     NOT NULL UNIQUE,
    name            VARCHAR(100)    NOT NULL,
    created_by      INTEGER     DEFAULT 0,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER     DEFAULT 0,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted         BOOLEAN     DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS province (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(20)     NOT NULL UNIQUE,
    name            VARCHAR(100)    NOT NULL,
    parent_id       INTEGER     DEFAULT 0,
    created_by      INTEGER     DEFAULT 0,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER     DEFAULT 0,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS district (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(20)     NOT NULL UNIQUE,
    name            VARCHAR(100)    NOT NULL,
    parent_id       INTEGER     DEFAULT 0,
    created_by      INTEGER     DEFAULT 0,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER     DEFAULT 0,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted         BOOLEAN     DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS sub_district (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(20)     NOT NULL UNIQUE,
    name            VARCHAR(100)    NOT NULL,
    parent_id       INTEGER     DEFAULT 0,
    created_by      INTEGER     DEFAULT 0,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER     DEFAULT 0,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted         BOOLEAN     DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS urban_village (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(20)     NOT NULL UNIQUE,
    name            VARCHAR(100)    NOT NULL,
    parent_id       INTEGER     DEFAULT 0,
    created_by      INTEGER     DEFAULT 0,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER     DEFAULT 0,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted         BOOLEAN     DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS company_profile (
    id              SERIAL PRIMARY KEY,
    npwp            VARCHAR(20)     NOT NULL UNIQUE,
    name            VARCHAR(200)    NOT NULL,
    address_1       VARCHAR(250)    DEFAULT '',
    address_2       VARCHAR(250)    DEFAULT '',
    country_id		INTEGER			DEFAULT 0,
    province_id		INTEGER			DEFAULT 0,
    district_id		INTEGER			DEFAULT 0,
    sub_district_id		INTEGER			DEFAULT 0,
    urban_village_id	INTEGER			DEFAULT 0,
    url_photo       VARCHAR(500)    DEFAULT '',
    created_by      INTEGER     DEFAULT 0,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER     DEFAULT 0,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted         BOOLEAN     DEFAULT FALSE
);
-- +migrate StatementEnd