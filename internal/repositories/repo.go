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

type User interface {
	GetUserById(id int) (interface{}, error)
	CreateUser(user models.UserCreate) (int, error)
	DeleteUser(id int) error
	UpdateUser(id int, user models.UserCreate) error
	GetAllUsers() (interface{}, error)
}


type Repo struct {
	Cinema
	User
}

func CinemaRepo(db *sqlx.DB) *Repo {
	return &Repo{
		Cinema: NewCinemaDB(db),
		User: NewuserDB(db),
	}
}
