package bookings

// CS stores the classes information
//
// should use mutex here, since it is a shared data
var bs BookingStore

// this line will ensure the ClassStore to implement the methods (compile time check)
var _ Bookings = &BookingStore{}

// AddBooking adds new booking to booking store
func (bs *BookingStore) AddBooking(booking Booking) int {
	bs.nextId++
	booking.ID = bs.nextId
	bs.bookings = append(bs.bookings, booking)

	return booking.ID
}
