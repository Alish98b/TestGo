package models

import "time"

type Ticket struct {
	ID       uint      `db:"id"`
	MovieID  uint      `db:"movie_id"`
	Seat     string    `db:"seat"`
	Price    float64   `db:"price"`
	ShowTime time.Time `db:"show_time"`
}

type TicketCreate struct {
	MovieID  uint
	Seat     string
	Price    float64
	ShowTime time.Time
}
