package services

import (
	"hotel/internal/models"
	"hotel/internal/repositories"
)

type Film interface {
	CreateFilm(film models.Film) (int, string)
}

type Service struct {
	Film
}

func NewServices(str *repositories.Storage) *Service {
	return &Service{
		Film: NewFilmService(str),
	}
}
