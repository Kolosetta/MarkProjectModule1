package user

import (
	"MarkProjectModule1/internal/models"
	"sync"
)

type Repository interface {
	Create(u models.User) error
	Get(username string) (models.User, error)
	GetList() ([]models.User, error)
}

type InMemoryRepository struct {
	mu        sync.RWMutex
	users     map[string]models.User
	currentId int64
}
