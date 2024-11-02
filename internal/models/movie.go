package models

import "time"

type Movie struct {
	ID          uint      `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ReleaseDate time.Time `db:"release_date"`
	Duration    int       `db:"duration"`
	Genres      []Genre   `db:"-"`
}

type Genre struct {
	ID     uint    `db:"id"`
	Name   string  `db:"name"`
	Movies []Movie `db:"-"`
}

type MovieCreate struct {
	Title       string
	Description string
	ReleaseDate time.Time
	Duration    int
}
