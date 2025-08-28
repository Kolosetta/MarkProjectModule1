package post

import (
	"MarkProjectModule1/internal/models"
	"MarkProjectModule1/pkg/events"
)

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
	events.LogEvent("Post Created", post)
	return nil
}

func (s *Service) GetPostList() []models.Post {
	var list, err = s.repo.GetList()
	if err != nil {
		return nil
	}
	return list
}

func (s *Service) LikePost(postId int64, userId int64) error {
	err := s.repo.LikePost(postId, userId)
	if err != nil {
		return err
	}
	events.LogEvent("Like", map[string]interface{}{
		"postId": postId,
		"userId": userId,
	})
	return nil
}
