-- +migrate Up
-- +migrate StatementBegin

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