package post

import (
	"MarkProjectModule1/internal/models"
)

type Repository interface {
	Create(post models.Post) error
	GetList() []models.Post
	LikePost(postId int64, userId int64) error
}

type InMemoryRepository struct {
	posts     []models.Post
	currentId int64
}

var repoInstance = &InMemoryRepository{
	posts: make([]models.Post, 0),
}

func GetRepository() *InMemoryRepository {
	return repoInstance
}

func (rep *InMemoryRepository) Create(post models.Post) error {
	rep.currentId++
	post.Id = rep.currentId
	rep.posts = append(rep.posts, post)
	return nil
}

func (rep *InMemoryRepository) GetList() []models.Post {
	return rep.posts
}

func (rep *InMemoryRepository) LikePost(postId int64, userId int64) error {
	postIndex := postId - 1
	post := rep.posts[postIndex]
	post.Likes = append(post.Likes, userId)
	rep.posts[postIndex] = post
	return nil
}
