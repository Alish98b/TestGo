package services

import (
	"hotel/internal/models"
	"hotel/internal/repositories"
)

type MovieService struct {
	repo *repositories.Repo
}

func NewMovieService(repo *repositories.Repo) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) GetMovieById(id int) (interface{}, error) {
	return s.repo.GetMovieById(id)
}

func (s *MovieService) CreateMovie(movie models.MovieCreate) (int, error) {
	return s.repo.CreateMovie(movie)
}

func (s *MovieService) DeleteMovie(id int) error {
	return s.repo.DeleteMovie(id)
}
func (s *MovieService) GetAllMovies() (interface{}, error) {
	return s.repo.GetAllMovies()
}

func (s *MovieService) UpdateMovie(id int, movie models.MovieCreate) error {
	return s.repo.UpdateMovie(id, movie)
}
