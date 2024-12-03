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

type Hall interface {
	GetHallById(id int) (models.Hall, error)
	CreateHall(hall models.HallCreate) (int, error)
	GetAllHalls() interface{}
	DeleteHall(id int) error
	UpdateHall(id int, hall models.Hall) error
}

type Session interface {
	GetSessionById(id int) (models.Session, error)
	CreateSession(session models.SessionCreate) (int, error)
	DeleteSession(id int) error
	UpdateSession(id int, session models.Session) error
	GetAllSessions() (models.Session, error)
}
type Repo struct {
	Cinema
	Hall
	Session
}

func CinemaRepo(db *sqlx.DB) *Repo {
	return &Repo{
		Cinema:  NewCinemaDB(db),
		Hall:    NewHallPostgres(db),
		Session: NewSessionPostgres(db),
	}
}
