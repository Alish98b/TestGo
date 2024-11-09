package models

import "time"

type User struct {
	ID       uint   `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
	//Bookings    []Booking   `db:"-"`
}

type Booking struct {
	ID        uint      `db:"id"`
	UserID    string    `db:"userid"`
	User      string    `db:"user"`
	CreatedAt time.Time `db:"createdat"`
	//Tickets		[]Ticket   `db:"-"`
}

type UserCreate struct {
	Name     string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

type UserLogin struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
