package post

import (
	"MarkProjectModule1/internal/models"
	"testing"
)

type mockRepository struct {
	createMockFunc   func(models.Post) error
	getListMockFunc  func() []models.Post
	likePostMockFunc func(postId int64, userId int64) error
}

func (m *mockRepository) Create(p models.Post) error {
	return m.createMockFunc(p)
}
func (m *mockRepository) GetList() []models.Post {
	return m.getListMockFunc()
}
func (m *mockRepository) LikePost(postId int64, userId int64) error {
	return m.likePostMockFunc(postId, userId)
}

func TestCreatePost_Success(t *testing.T) {
	// Мокаем репозиторий
	repo := &mockRepository{
		createMockFunc: func(p models.Post) error {
			if p.Author != "Igor" || p.Text != "His Text" {
				t.Errorf("unexpected post: %+v", p)
			}
			return nil
		},
	}

	svc := NewService(repo)
	err := svc.CreatePost(models.Post{
		Author: "Igor",
		Text:   "His Text",
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetPostList_Success(t *testing.T) {
	expected := []models.Post{
		{Id: 1, Author: "A"},
		{Id: 2, Author: "B"},
	}

	repo := &mockRepository{
		getListMockFunc: func() []models.Post {
			return expected
		},
	}

	svc := NewService(repo)
	result := svc.GetPostList()

	if len(result) != len(expected) {
		t.Errorf("expected %d posts, got %d", len(expected), len(result))
	}
}

func TestLikePost_Success(t *testing.T) {
	repo := &mockRepository{
		likePostMockFunc: func(postId int64, userId int64) error {
			if postId != 5 || userId != 10 {
				t.Errorf("unexpected like params: %d, %d", postId, userId)
			}
			return nil
		},
	}

	svc := NewService(repo)
	err := svc.LikePost(5, 10)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
