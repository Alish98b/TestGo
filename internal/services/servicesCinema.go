package services

import (
	"hotel/internal/models"
	"hotel/internal/repositories"
)

type Movie interface {
	GetMovieById(id int) (interface{}, error)
	CreateMovie(movie models.MovieCreate) (int, error)
	DeleteMovie(id int) error
	UpdateMovie(id int, movie models.MovieCreate) error
	GetAllMovies() (interface{}, error)
}

type Genre interface {
	CreateGenre(genre models.GenreCreate) (int, error)
	GetGenreById(id int) (models.Genre, error)
	GetAllGenres() ([]models.Genre, error)
	UpdateGenre(id int, genre models.GenreCreate) error
	DeleteGenre(id int) error
}

type Ticket interface {
	CreateTicket(ticket models.TicketCreate) (int, error)
	GetTicketById(id int) (models.Ticket, error)
	GetAllTickets() ([]models.Ticket, error)
	UpdateTicket(id int, ticket models.TicketCreate) error
	DeleteTicket(id int) error
}

type ServicesCinema struct {
	Movie
	Genre
	Ticket
}

func NewServicesCinema(repo *repositories.Repo) *ServicesCinema {
	return &ServicesCinema{
		Movie:  NewMovieService(repo),
		Genre:  NewGenreService(repo),
		Ticket: NewTicketService(repo),
	}
}
