package user

import (
	"MarkProjectModule1/internal/models"
	"testing"
)

type mockRepository struct {
	createMockFunc  func(u models.User) error
	GetMockFunc     func(username string) (models.User, error)
	GetListMockFunc func() ([]models.User, error)
}

func (m *mockRepository) Create(u models.User) error {
	return m.createMockFunc(u)
}
func (m *mockRepository) Get(username string) (models.User, error) {
	return m.GetMockFunc(username)
}
func (m *mockRepository) GetList() ([]models.User, error) {
	return m.GetListMockFunc()
}

func TestCreateUser_Success(t *testing.T) {
	// Мокаем репозиторий
	repo := &mockRepository{
		createMockFunc: func(u models.User) error {
			if u.Username != "Igor222" || u.Email != "igor@mail.ru" {
				t.Errorf("unexpected post: %+v", u)
			}
			return nil
		},
	}

	service := NewService(repo)
	err := service.CreateUser("Igor222", "igor@mail.ru")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetList_Success(t *testing.T) {
	// Мокаем репозиторий
	repo := &mockRepository{
		GetListMockFunc: func() ([]models.User, error) {
			return []models.User{
				{Username: "Igor222", Email: "igor@mail.ru"},
				{Username: "Ivan", Email: "ivan@gmail.com"},
			}, nil
		},
	}

	service := NewService(repo)
	users, err := service.GetUsersList()

	if err != nil {
		t.Error("unexpected error")
	}

	if len(users) != 2 {
		t.Errorf("unexpected users: %+v", users)
	}
}

func TestGetUser_Success(t *testing.T) {
	// Мокаем репозиторий
	repo := &mockRepository{
		GetMockFunc: func(username string) (models.User, error) {
			return models.User{
				Id:       1,
				Username: "Igor222",
				Email:    "igor@mail.ru",
			}, nil
		},
	}

	service := NewService(repo)
	user, err := service.GetUser("Igor222")

	if user.Username != "Igor222" || user.Email != "igor@mail.ru" {
		t.Errorf("unexpected user: %+v", user)
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
