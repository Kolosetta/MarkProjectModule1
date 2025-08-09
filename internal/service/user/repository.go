package user

import (
	"MarkProjectModule1/internal/models"
	"errors"
)

type Repository interface {
	Create(u models.User) error
	Get(username string) (models.User, error)
	GetList() []models.User
}

type InMemoryRepository struct {
	users     map[string]models.User
	currentId int64
}

var repoInstance = &InMemoryRepository{
	users: make(map[string]models.User),
}

func GetRepository() *InMemoryRepository {
	return repoInstance
}

func (rep *InMemoryRepository) Create(user models.User) error {
	if _, exists := rep.users[user.Username]; exists {
		return errors.New("user already exists")
	}

	rep.currentId++
	user.Id = rep.currentId
	rep.users[user.Username] = user
	return nil
}

func (rep *InMemoryRepository) Get(username string) (models.User, error) {
	if user, exists := rep.users[username]; exists {
		return user, nil
	}
	return models.User{}, errors.New("user not found")
}

func (rep *InMemoryRepository) GetList() []models.User {
	result := make([]models.User, len(rep.users))
	for _, user := range rep.users {
		result = append(result, user)
	}
	return result
}
