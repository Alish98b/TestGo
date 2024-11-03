package services

import (
	"hotel/internal/models"
	"hotel/internal/repositories"
)

type UserService struct {
	repo *repositories.Repo
}

func NewUserService(repo *repositories.Repo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserById(id int) (interface{}, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService) CreateUser(user models.UserCreate) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
func (s *UserService) GetAllUsers() (interface{}, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) UpdateUser(id int, user models.UserCreate) error {
	return s.repo.UpdateUser(id, user)
}
