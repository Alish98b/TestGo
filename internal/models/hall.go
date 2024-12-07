package models

type Hall struct {
	Id       uint   `db:"id"`
	Name     string `db:"name"`
	Capacity int    `db:"capacity"`
}

type HallCreate struct {
	Name     string `db:"name"`
	Capacity int    `db:"capacity"`
}
