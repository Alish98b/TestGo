package models

type Genre struct {
	ID     uint    `db:"id"`
	Name   string  `db:"name"`
	Movies []Movie `db:"-"`
}

type GenreCreate struct {
	Name *string
}
