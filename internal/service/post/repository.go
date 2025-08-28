package post

import (
	"MarkProjectModule1/internal/models"
)

type Repository interface {
	Create(post models.Post) error
	GetList() ([]models.Post, error)
	LikePost(postId int64, userId int64) error
}
