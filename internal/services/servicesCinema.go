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

type ServicesCinema struct {
	Movie
}

func NewServicesCinema(repo *repositories.Repo) *ServicesCinema {
	return &ServicesCinema{
		Movie: NewMovieService(repo),
	}
}
