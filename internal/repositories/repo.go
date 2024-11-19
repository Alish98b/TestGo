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

type Genre interface {
    CreateGenre(genre models.GenreCreate) (int, error)
    GetGenreById(id int) (models.Genre, error)
    GetAllGenres() ([]models.Genre, error)
    UpdateGenre(id int, genre models.GenreCreate) error
    DeleteGenre(id int) error
}

type Repo struct {
	Cinema
	Genre
}

func CinemaRepo(db *sqlx.DB) *Repo {
	return &Repo{
		Cinema: NewCinemaDB(db),
		Genre:  NewGenreDB(db),
	}
}
