package services

import (
	"hotel/internal/models"
	"hotel/internal/repositories"
)

type Film interface {
	GetFilmById(id int) (models.Film, error)
	CreateFilm(room models.Film) (int, string)
	DeleteFilm(id int) (int, string)
	UpdateFilm(film models.Film) (interface{}, string)
	GetAllFilms() interface{}
}

type Service struct {
	Film
}

func NewServices(str *repositories.Storage) *Service {
	return &Service{
		Film: NewFilmService(str),
	}
}
