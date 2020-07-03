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

func (s Service) BookPropertyById(propertyId int, userId int, startDate time.Time, endDate time.Time) (User, Booking, Property, error) {
	insertBooking := `
	INSERT INTO bookings.durations(property_id, user_id, start_date, end_date)
	VALUES
		(:property_id, :user_id, :start_date, :end_date)
	RETURNING *;
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

	err = trans.Commit()
	if err != nil {
		return User{}, Booking{}, Property{}, err
	}

	property, err = s.GetPropertyByID(propertyId)
	if err != nil {
		return User{}, Booking{}, Property{}, err
	}

	user, err = s.GetUserByID(userId)
	if err != nil {
		return User{}, Booking{}, Property{}, err
	}

	return user, booking, property, nil
}

func (s Service) GetPropertyByID(propertyId int) (Property, error) {
	type propByID struct {
		ID int `db:"id"`
	}

	propertyQuery := `
	SELECT * FROM bookings.properties WHERE id = :id;
	`
	var property Property

	trans, err := s.db.Beginx()
	if err != nil {
		return Property{}, err
	}

	defer trans.Rollback() //nolint:errcheck

	stmt, err := trans.PrepareNamed(propertyQuery)

	if err != nil {
		return Property{}, err
	}

	tempProperty := propByID{
		ID: propertyId,
	}

	err = stmt.Get(&property, tempProperty)
	if err != nil {
		return Property{}, err
	}

	err = trans.Commit()
	if err != nil {
		return Property{}, err
	}

	return property, nil
}

func (s Service) GetUserByID(userId int) (User, error) {
	type userByID struct {
		ID int `db:"id"`
	}

	userQuery := `
	SELECT * FROM bookings.users WHERE id = :id;
	`
	var user User

	trans, err := s.db.Beginx()
	if err != nil {
		return User{}, err
	}

	defer trans.Rollback() //nolint:errcheck

	stmt, err := trans.PrepareNamed(userQuery)
	if err != nil {
		return User{}, err
	}

	tempUser := userByID{
		ID: userId,
	}

	err = stmt.Get(&user, tempUser)
	if err != nil {
		return User{}, err
	}

	err = trans.Commit()
	if err != nil {
		return User{}, err
	}

	return user, nil
}
