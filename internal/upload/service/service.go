package service

import (
	"ortak/internal/upload/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SaveFile(file repository.File) error {
	return s.repo.Create(file)
}

func (s *Service) GetFile(id string) (*repository.File, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetUserFiles(userID string) ([]repository.File, error) {
	return s.repo.GetByUserID(userID)
}

func (s *Service) GetAllFiles() ([]repository.File, error) {
	return s.repo.GetAll()
}
