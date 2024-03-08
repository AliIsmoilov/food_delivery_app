CREATE TABLE IF NOT EXISTS couriers(
    id UUID PRIMARY KEY,
    name VARCHAR(256),
    phone_number VARCHAR(24) UNIQUE NOT NULL,
    photo VARCHAR,
    car_model VARCHAR(256),
    car_number VARCHAR(56),
    car_color VARCHAR(128),
    is_available BOOLEAN DEFAULT TRUE,
    latitude DECIMAL(7,5),
    longitude DECIMAL(8,5),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);