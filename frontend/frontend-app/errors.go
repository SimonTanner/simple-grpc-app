package frontend

import (
	"errors"
)

var (
	errNoUserID     = errors.New("no user ID in params")
	errNoBookingID  = errors.New("no booking ID in params")
	errNoPropertyID = errors.New("no property ID in params")
	errNoBookings   = errors.New("no bookings returned from backend")
)
