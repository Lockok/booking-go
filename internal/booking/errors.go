package booking

import "errors"

var (
	ErrSeatAlreadyBooked = errors.New("seat is already taken")
)