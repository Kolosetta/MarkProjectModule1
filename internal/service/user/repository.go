package user

import (
	"MarkProjectModule1/internal/models"
	"errors"
	"sync"
)

type Repository interface {
	Create(u models.User) error
	Get(username string) (models.User, error)
	GetList() []models.User
}

type InMemoryRepository struct {
	mu        sync.RWMutex
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
	rep.mu.Lock()
	defer rep.mu.Unlock()
	if _, exists := rep.users[user.Username]; exists {
		return errors.New("user already exists")
	}

	rep.currentId++
	user.Id = rep.currentId
	rep.users[user.Username] = user
	return nil
}

func (rep *InMemoryRepository) Get(username string) (models.User, error) {
	rep.mu.Lock()
	defer rep.mu.Unlock()
	if user, exists := rep.users[username]; exists {
		return user, nil
	}
	return models.User{}, errors.New("user not found")
}

func (rep *InMemoryRepository) GetList() []models.User {
	rep.mu.Lock()
	defer rep.mu.Unlock()
	result := make([]models.User, len(rep.users))
	for _, user := range rep.users {
		result = append(result, user)
	}
	return result
}
