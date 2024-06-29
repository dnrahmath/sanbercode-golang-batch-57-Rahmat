-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE person {
    id BIGINT NOT NULL,
    first_name VARCHAR(256),
    last_name VARCHAR(256)
}

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    title VARCHAR(256) NOT NULL,
    description TEXT,
    image_url VARCHAR(512),
    release_year INT,
    price VARCHAR(32),
    total_page INT,
    thickness VARCHAR(32),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    category_id INT REFERENCES category(id)
);

-- +migrate StatementEnd