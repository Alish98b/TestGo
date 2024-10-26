package services

import (
	"hotel/internal/models"
	"hotel/internal/repositories"
)

type RoomService struct {
	storage *repositories.Storage
}

func NewRoomService(str *repositories.Storage) *RoomService {
	return &RoomService{storage: str}
}

func (s *RoomService) GetRoomById(id int) (models.Room, error) {
	return s.storage.GetRoomById(id)
}

func (s *RoomService) CreateFilm(film models.Film) (int, string) {
	return s.storage.CreateRoom(film)
}
