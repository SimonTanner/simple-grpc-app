package bookings

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

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

func (s Service) BookPropertyById(propertyId, userId int32, startDate, endDate time.Time) (User, Booking, Property, error) {
	insertBooking := `
	INSERT INTO bookings.durations(property_id, user_id, start_date, end_date)
	VALUES
		(:property_id, :user_id, :start_date, :end_date)
	RETURNING *;
	`

	userQuery := `
	SELECT * FROM bookings.users WHERE id = :id;
	`

	propertyQuery := `
	SELECT * FROM bookings.properties WHERE id = :id;
	`
	var (
		booking  Booking
		user     User
		property Property
	)

	newBooking := Booking{
		PropertyId: propertyId,
		UserId:     userId,
		StartDate:  startDate,
		EndDate:    endDate,
	}

	trans, err := s.db.Beginx()
	if err != nil {
		return User{}, Booking{}, Property{}, err
	}

	defer trans.Rollback() //nolint:errcheck

	stmt, err := trans.PrepareNamed(insertBooking)
	if err != nil {
		return User{}, Booking{}, Property{}, err
	}

	defer stmt.Close()

	err = stmt.Get(&booking, &newBooking)
	if err != nil {
		return User{}, Booking{}, Property{}, err
	}

	stmt, err = trans.PrepareNamed(userQuery)
	if err != nil {
		return User{}, Booking{}, Property{}, err
	}

	err = stmt.Select(&user, &userId)
	if err != nil {
		return User{}, Booking{}, Property{}, err
	}

	stmt, err = trans.PrepareNamed(propertyQuery)
	if err != nil {
		return User{}, Booking{}, Property{}, err
	}

	err = stmt.Select(&property, &propertyId)
	if err != nil {
		return User{}, Booking{}, Property{}, err
	}

	err = trans.Commit()
	if err != nil {
		return User{}, Booking{}, Property{}, err
	}

	return user, booking, property, nil
}
