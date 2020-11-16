BEGIN;

-- DROP TABLE schema_migrations;

CREATE SCHEMA IF NOT EXISTS bookings;

CREATE TABLE IF NOT EXISTS bookings.users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS bookings.properties (
    id SERIAL PRIMARY KEY,
    door_number VARCHAR(50),
    address VARCHAR(1000),
    city VARCHAR(100),
    country VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS bookings.durations (
    id SERIAL PRIMARY KEY,
    property_id INTEGER NOT NULL REFERENCES bookings.properties(id),
    user_id INTEGER NOT NULL REFERENCES bookings.users(id),
    start_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    end_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMIT;
