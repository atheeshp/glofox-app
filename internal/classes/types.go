package classes

import "time"

// Bookings interface which will give us more flexibility rewrite the logics in the future
type Classes interface {
	AddClass(class Class) int
}

type reqClass struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

type Class struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Capacity  int       `json:"capacity"`
}

func NewClass(name string, start, end time.Time, capacity int) Class {
	return Class{
		Name:      name,
		StartDate: start,
		EndDate:   end,
		Capacity:  capacity,
	}
}

type ClassStore struct {
	classes []Class
	nextID  int
}
