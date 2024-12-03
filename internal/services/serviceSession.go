package services

import (
	"hotel/internal/models"
	_ "hotel/internal/models"
	"hotel/internal/repositories"
)

type SessionService struct {
	repo *repositories.Repo
}

func NewSessionService(repo *repositories.Repo) *SessionService {
	return &SessionService{repo: repo}
}

func (s *SessionService) CreateSession(session models.SessionCreate) (int, error) {
	return s.repo.CreateSession(session)
}

func (s *SessionService) DeleteSession(id int) error {
	return s.repo.DeleteSession(id)
}

func (s *SessionService) GetSessionById(id int) (models.Session, error) {
	return s.repo.GetSessionById(id)
}

func (s *SessionService) UpdateSession(id int, session models.Session) error {
	return s.repo.UpdateSession(id, session)
}

func (s *SessionService) GetAllSessions() (models.Session, error) {
	return s.repo.GetAllSessions()
}
