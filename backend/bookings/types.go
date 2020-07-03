package bookings

import "time"

type Property struct {
	Id         int32     `db:"id"`
	DoorNumber string    `db:"door_number"`
	Address    string    `db:"address"`
	City       string    `db:"city"`
	Country    string    `db:"country"`
	CreatedAt  time.Time `db:"created_at"`
}

type User struct {
	Id        int32     `db:"id"`
	FirstName string    `db:"first_name"`
	Surname   string    `db:"surname"`
	CreatedAt time.Time `db:"created_at"`
}

type Booking struct {
	PropertyId int32     `db:"property_id"`
	UserId     int32     `db:"user_id"`
	StartDate  time.Time `db:"start_date"`
	EndDate    time.Time `db:"end_date"`
	CreatedAt  time.Time `db:"created_at"`
}
