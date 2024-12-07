package services

import (
	"hotel/internal/models"
	"hotel/internal/repositories"
)

type GenreService struct {
	repo *repositories.Repo
}

func NewGenreService(repo *repositories.Repo) *GenreService {
	return &GenreService{repo: repo}
}

func (s *GenreService) CreateGenre(genre models.GenreCreate) (int, error) {
	// err := fmt.Errorf("Name %s", "is required")
	// if genre.Name == nil {
	// 	return 0, err
	// }
	return s.repo.Genre.CreateGenre(genre)
}

func (s *GenreService) GetGenreById(id int) (models.Genre, error) {
	return s.repo.Genre.GetGenreById(id)
}

func (s *GenreService) GetAllGenres() ([]models.Genre, error) {
	return s.repo.Genre.GetAllGenres()
}

func (s *GenreService) UpdateGenre(id int, genre models.GenreCreate) error {
	return s.repo.Genre.UpdateGenre(id, genre)
}

func (s *GenreService) DeleteGenre(id int) error {
	return s.repo.Genre.DeleteGenre(id)
}
