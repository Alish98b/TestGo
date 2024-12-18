package services

import (
	"hotel/internal/models"
	"hotel/internal/repositories"
)

type HallService struct {
	repo *repositories.Repo
}

func NewHallService(repo *repositories.Repo) *HallService {
	return &HallService{repo: repo}
}

func (s *HallService) GetHallById(id int) (models.Hall, error) {
	return s.repo.GetHallById(id)
}

func (s *HallService) CreateHall(hall models.HallCreate) (int, error) {
	return s.repo.CreateHall(hall)
}

func (s *HallService) DeleteHall(id int) error {
	return s.repo.DeleteHall(id)
}

func (s *HallService) UpdateHall(id int, hall models.Hall) error {
	return s.repo.UpdateHall(id, hall)
}

func (s *HallService) GetAllHalls() interface{} {
	return s.repo.GetAllHalls()
}

//func (s *HallService) GetHallByName(name string) (*models.Hall, error) {
//	var hall models.Hall
//	err := s.db.Where("name = ?", name).First(&hall).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, nil // Зал с таким именем не найден
//		}
//		return nil, err // Другая ошибка
//	}
//	return &hall, nil
//}
