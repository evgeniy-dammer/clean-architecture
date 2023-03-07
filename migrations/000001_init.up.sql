-- +goose Up
-- +goose StatementBegin

CREATE SCHEMA IF NOT EXISTS clean;

CREATE TABLE IF NOT EXISTS clean.contact
(
    id UUID DEFAULT gen_random_uuid() NOT NULL CONSTRAINT pk_contact PRIMARY KEY,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    name VARCHAR(50) DEFAULT '':: CHARACTER VARYING NOT NULL,
    surname VARCHAR(100) DEFAULT '':: CHARACTER VARYING NOT NULL,
    patronymic VARCHAR(100) DEFAULT '':: CHARACTER VARYING NOT NULL,
    email VARCHAR(250),
    phone_number VARCHAR(50),
    age SMALLINT CONSTRAINT age_check CHECK ((age >= 0) AND (age <= 200)),
    gender SMALLINT,
    is_archived BOOLEAN DEFAULT FALSE NOT NULL
);

CREATE TABLE IF NOT EXISTS clean."group"
(
    id UUID DEFAULT gen_random_uuid() NOT NULL CONSTRAINT pk_group PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(250) NOT NULL,
    description VARCHAR(1000),
    contact_count bigint DEFAULT 0 NOT NULL,
    is_archived BOOLEAN DEFAULT FALSE NOT NULL
);

CREATE TABLE IF NOT EXISTS clean.contact_in_group
(
    id UUID DEFAULT gen_random_uuid() NOT NULL CONSTRAINT pk_contact_in_group PRIMARY KEY,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    modified_at timestamp DEFAULT CURRENT_TIMESTAMP,
    contact_id UUID NOT NULL CONSTRAINT fk_contact_id REFERENCES clean.contact,
    group_id UUID NOT NULL CONSTRAINT fk_group_id REFERENCES clean."group"
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS clean.contact_in_group;
DROP TABLE IF EXISTS clean.contact;
DROP TABLE IF EXISTS clean.group;
DROP SCHEMA IF EXISTS clean;


-- +goose StatementEnd