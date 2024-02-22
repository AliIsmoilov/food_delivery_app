CREATE TABLE IF NOT EXISTS restaurants (
    id uuid PRIMARY KEY,
    name VARCHAR(64),
    adress VARCHAR(64),
    phone_number VARCHAR(16) UNIQUE,
    photo VARCHAR,
    latitude DECIMAL(7,5) NOT NULL,
    longitude DECIMAL(7,5) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);