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
