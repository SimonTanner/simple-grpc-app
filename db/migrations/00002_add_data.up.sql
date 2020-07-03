BEGIN;

INSERT INTO bookings.properties(door_number, address, city, country) VALUES
('87', 'Cadogan Terrace', 'London', 'UK'),
('104', 'Ladbroke Grove', 'London', 'UK'),
('26', 'Sussex Gardens', 'London', 'UK'),
('6b', 'Rue Saint Martin', 'Paris', 'France'),
('34', 'Boulevard Voltaire', 'Paris', 'France'),
('16', 'Calle Menorca', 'Madrid', 'Spain'),
('78', 'Travessera de Gracia', 'Barcelona', 'Spain'),
('19', 'Calle Rio Seco', 'Alicante', 'Spain'),
('41c', 'Calle de Padilla', 'Madrid', 'Spain'),
('23a', 'Calle de Padilla', 'Madrid', 'Spain');



INSERT INTO bookings.users(first_name, surname) VALUES
('Montgomery', 'Burns'),
('Billy-Bob', 'Thornton');

COMMIT;