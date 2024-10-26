package services

import (
	"hotel/internal/models"
	"hotel/internal/repositories"
)

type FilmService struct {
	storage *repositories.Storage
}

func NewFilmService(str *repositories.Storage) *FilmService {
	return &FilmService{storage: str}
}

func (s *FilmService) CreateFilm(film models.Film) (int, string) {
	return s.storage.CreateFilm(film)
}
