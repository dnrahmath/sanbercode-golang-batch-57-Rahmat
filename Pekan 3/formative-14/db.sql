CREATE DATABASE "db-go-sql";

-- Gunakan \cuntuk beralih ke database yang baru dibuat
\c "db-go-sql"

-- Buat tabel di dalam database "db-go-sql"
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(50) NOT NULL,
    email TEXT UNIQUE NOT NULL,
    age INT NOT NULL,
    division VARCHAR(20) NOT NULL
);