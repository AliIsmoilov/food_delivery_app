DROP TYPE IF EXISTS gender;
CREATE TYPE gender AS ENUM('male', 'female');

CREATE TABLE IF NOT EXISTS customers (
    id uuid PRIMARY KEY NOT NULL,
    email VARCHAR(255),
    phone_number VARCHAR(16) UNIQUE,
    name VARCHAR(64),
    photo VARCHAR(255),
    sex gender,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

