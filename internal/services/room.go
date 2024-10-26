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

func (s *FilmService) GetFilmById(id int) (models.Film, error) {
	return s.storage.GetFilmById(id)
}

func (s *FilmService) CreateFilm(film models.Film) (int, string) {
	return s.storage.CreateFilm(film)
}

func (s *FilmService) DeleteFilm(id int) (int, string) {
	return s.storage.DeleteFilm(id)
}

func (s *FilmService) UpdateFilm(film models.Film) (interface{}, string) {
	return s.storage.UpdateFilm(film)
}

func (s *FilmService) GetAllFilms() interface{} {
	return s.storage.GetAllFilms()
}
