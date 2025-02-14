package bookings

import "time"

type Bookings interface {
	AddBooking(booking Booking) int
}

type reqBooking struct {
	Member string `json:"member"`
	Date   string `json:"date"`
}

type Booking struct {
	ID     int       `json:"id"`
	Member string    `json:"member"`
	Date   time.Time `json:"date"`
}

func NewBooking(member string, date time.Time) Booking {
	return Booking{
		Member: member,
		Date:   date,
	}
}

type BookingStore struct {
	bookings []Booking
	nextId   int
}
