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
type Hall interface {
	GetHallById(id int) (models.Hall, error)
	CreateHall(hall models.HallCreate) (int, error)
	GetAllHalls() interface{}
	DeleteHall(id int) (int error)
	UpdateHall(id int, hall models.Hall) error
}

type Session interface {
	GetSessionById(id int) (models.Session, error)
	CreateSession(session models.SessionCreate) (int, error)
	DeleteSession(id int) error
	UpdateSession(id int, session models.Session) error
	GetAllSessions() (models.Session, error)
}
type Genre interface {
	CreateGenre(genre models.GenreCreate) (int, error)
	GetGenreById(id int) (models.Genre, error)
	GetAllGenres() ([]models.Genre, error)
	UpdateGenre(id int, genre models.GenreCreate) error
	DeleteGenre(id int) error
}
type ServicesCinema struct {
	Movie
	Hall
	Session
	Genre
}

func NewServicesCinema(repo *repositories.Repo) *ServicesCinema {
	return &ServicesCinema{
		Movie:   NewMovieService(repo),
		Hall:    NewHallService(repo),
		Session: NewSessionService(repo),
		Genre:   NewGenreService(repo),
	}
}
