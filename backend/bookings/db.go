package bookings

import "github.com/jmoiron/sqlx"

type Service struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Service {
	return Service{
		db: db,
	}
}

func (s Service) GetAllProperties() ([]Property, error) {
	var properties []Property

	const query = `
	SELECT
		id,
		door_number,
		address,
		city,
		country,
		created_at
	FROM
		bookings.properties;
	`

	err := s.db.Select(&properties, query)
	if err != nil {
		return nil, err
	}

	return properties, nil
}
