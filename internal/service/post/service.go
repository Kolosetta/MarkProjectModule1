package post

import "MarkProjectModule1/internal/models"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePost(post models.Post) error {
	if err := s.repo.Create(post); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetPostList() []models.Post {
	return s.repo.GetList()
}

func (s *Service) LikePost(postId int64, userId int64) error {
	err := s.repo.LikePost(postId, userId)
	if err != nil {
		return err
	}
	return nil
}
