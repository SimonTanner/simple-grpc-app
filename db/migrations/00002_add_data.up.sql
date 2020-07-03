BEGIN;

INSERT INTO bookings.properties(door_number, address, city, country) VALUES
('87', 'Cadogan Terrace', 'London', 'UK'),
('6b', 'Rue Saint Martin', 'Paris', 'France');

INSERT INTO bookings.users(first_name, surname) VALUES
('Montgomery', 'Burns'),
('Billy-Bob', 'Thornton');

COMMIT;