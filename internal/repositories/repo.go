package repositories

import (
	"hotel/internal/models"

	"github.com/jmoiron/sqlx"
)

type Cinema interface {
	GetMovieById(id int) (interface{}, error)
	CreateMovie(movie models.MovieCreate) (int, error)
	UpdateMovie(id int, movie models.MovieCreate) error
	DeleteMovie(id int) error
	GetAllMovies() (interface{}, error)
}

type Repo struct {
	Cinema
}

func CinemaRepo(db *sqlx.DB) *Repo {
	return &Repo{
		Cinema: NewCinemaDB(db),
	}
}
