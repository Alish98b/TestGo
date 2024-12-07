package models

import "time"

type Session struct {
	ID        uint      `db:"primaryKey"`
	MovieID   uint      `db:"not null"`
	Movie     Movie     `db:"foreignKey:MovieID"`
	HallID    uint      `db:"not null"`
	Hall      Hall      `db:"foreignKey:HallID"`
	StartTime time.Time `db:"not null"`
	//Tickets   []Ticket  `db:"foreignKey:SessionID"`

}

type SessionCreate struct {
	MovieID   uint      `db:"movie_id"`
	HallID    uint      `db:"hall_id"`
	StartTime time.Time `db:"start_time"`
}
