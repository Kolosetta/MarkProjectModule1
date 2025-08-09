package user

import "MarkProjectModule1/internal/models"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(username string, email string) error {
	user := models.User{
		Username: username,
		Email:    email,
	}
	if err := s.repo.Create(user); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUser(username string) (models.User, error) {
	return s.repo.Get(username)
}

func (s *Service) GetUsersList() []models.User {
	return s.repo.GetList()

}
