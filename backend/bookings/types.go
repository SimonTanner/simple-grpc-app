package bookings

type Property struct {
	Id         int32  `db:"id"`
	DoorNumber string `db:"door_number"`
	Address    string `db:"address"`
	City       string `db:"city"`
	Country    string `db:"country"`
}
