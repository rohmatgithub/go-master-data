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

CREATE TABLE product_category (
	id              SERIAL PRIMARY KEY,
	code            varchar(20) NOT NULL,
	"name"          varchar(200) NOT NULL,
	company_id      INTEGER NOT NULL,
	created_by      INTEGER NULL DEFAULT 0,
	created_at      timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by      INTEGER NULL DEFAULT 0,
	updated_at      timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	deleted bool    NULL DEFAULT false,
	CONSTRAINT uq_productcategory_codecompanydivision UNIQUE (code, company_id, division_id)
);

CREATE TABLE product_group_hierarchy (
	id              SERIAL PRIMARY KEY,
	code            varchar(20) NOT NULL,
	company_id      INTEGER NOT NULL,
	"name"          varchar(200) NOT NULL,
	"level"         INTEGER NULL DEFAULT 0,
	parent_id       INTEGER NOT NULL DEFAULT 0,
	created_by      INTEGER NULL DEFAULT 0,
	created_at      timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by      INTEGER NULL DEFAULT 0,
	updated_at      timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	deleted bool    NULL DEFAULT false,
	CONSTRAINT uq_pgh_codecompanydivisionparent UNIQUE (code, company_id, division_id, parent_id)
);

CREATE TABLE product (
	id              SERIAL PRIMARY KEY,
	code            varchar(20) NOT NULL,
	company_id      INTEGER NOT NULL,
    category_id     INTEGER NOT NULL,
    group_id        INTEGER NOT NULL,
	"name"          varchar(200) NOT NULL DEFAULT '',
    selling_price   FLOAT NOT NULL DEFAULT 0.0,
    buying_price    FLOAT NOT NULL DEFAULT 0.0,
    uom_1           VARCHAR(10) NOT NULL DEFAULT '',
    uom_2           VARCHAR(10) NOT NULL DEFAULT '',
	conv_1_to_2     INTEGER NOT NULL DEFAULT 0,
	created_by      INTEGER NULL DEFAULT 0,
	created_at      timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by      INTEGER NULL DEFAULT 0,
	updated_at      timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	deleted bool    NULL DEFAULT false,
	CONSTRAINT uq_prod_companyidcode UNIQUE (company_id, code)
);

CREATE TABLE customer (
	id              SERIAL PRIMARY KEY,
	code            varchar(20) NOT NULL,
	company_id      INTEGER NOT NULL,
	"name"          varchar(200) NOT NULL DEFAULT '',
    phone		   	VARCHAR(20) DEFAULT '',
    email		   	VARCHAR(20) DEFAULT '',
    country_id		INTEGER	DEFAULT 0,
    province_id		INTEGER	DEFAULT 0,
    district_id		INTEGER	DEFAULT 0,
    sub_district_id		INTEGER	DEFAULT 0,
    urban_village_id	INTEGER	DEFAULT 0,
    address 		VARCHAR(200) DEFAULT 0,
	created_by      INTEGER NULL DEFAULT 0,
	created_at      timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by      INTEGER NULL DEFAULT 0,
	updated_at      timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	deleted bool    NULL DEFAULT false,
	CONSTRAINT uq_cust_companyidbranchidcode UNIQUE (company_id, branch_id, code)
);
-- +migrate StatementEnd