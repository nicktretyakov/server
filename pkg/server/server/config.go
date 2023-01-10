package server

import "be/pkg/server/booking"

type Config struct {
	Secret string
}

func (c Config) Booking() booking.Config {
	return booking.Config{}
}
