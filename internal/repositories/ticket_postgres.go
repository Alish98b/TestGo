package repositories

import (
	"hotel/internal/models"

	"github.com/jmoiron/sqlx"
)

type TicketDB struct {
	db *sqlx.DB
}

func NewTicketDB(db *sqlx.DB) *TicketDB {
	return &TicketDB{db: db}
}

func (t *TicketDB) CreateTicket(ticket models.TicketCreate) (int, error) {
	var id int
	query := `INSERT INTO tickets (movie_id, seat, price, show_time) VALUES ($1, $2, $3, $4) RETURNING id`
	row := t.db.QueryRow(query, ticket.MovieID, ticket.Seat, ticket.Price, ticket.ShowTime)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (t *TicketDB) GetTicketById(id int) (models.Ticket, error) {
	var ticket models.Ticket
	query := `SELECT id, movie_id, seat, price, show_time FROM tickets WHERE id=$1`
	err := t.db.Get(&ticket, query, id)
	return ticket, err
}

func (t *TicketDB) GetAllTickets() ([]models.Ticket, error) {
	var tickets []models.Ticket
	query := `SELECT id, movie_id, seat, price, show_time FROM tickets`
	err := t.db.Select(&tickets, query)
	return tickets, err
}

func (t *TicketDB) UpdateTicket(id int, ticket models.TicketCreate) error {
	query := `UPDATE tickets SET movie_id=$1, seat=$2, price=$3, show_time=$4 WHERE id=$5`
	_, err := t.db.Exec(query, ticket.MovieID, ticket.Seat, ticket.Price, ticket.ShowTime, id)
	return err
}

func (t *TicketDB) DeleteTicket(id int) error {
	query := `DELETE FROM tickets WHERE id=$1`
	_, err := t.db.Exec(query, id)
	return err
}
