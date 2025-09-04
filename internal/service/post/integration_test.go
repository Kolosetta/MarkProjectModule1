package post

import (
	"MarkProjectModule1/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"os"
	"testing"
)

var (
	pool    *pgxpool.Pool
	repo    *PostgresRepository
	service *Service
)

func TestMain(m *testing.M) {
	// Создаем докер пул
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=my_pass",
			"POSTGRES_DB=testdb",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// Удалить контейнер после завершения
	defer func() {
		_ = dockerPool.Purge(resource)
	}()

	// Ждем пока Postgres поднимется
	var dbpool *pgxpool.Pool
	if err := dockerPool.Retry(func() error {
		dsn := fmt.Sprintf("postgres://postgres:my_pass@localhost:%s/testdb?sslmode=disable", resource.GetPort("5432/tcp"))
		dbpool, err = pgxpool.New(context.Background(), dsn)
		if err != nil {
			return err
		}
		return dbpool.Ping(context.Background())
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	// Сохраняем глобально
	pool = dbpool
	repo = NewPostgresRepository(pool)
	service = NewService(repo)

	// Создаем таблицы
	_, err = pool.Exec(context.Background(), `
	CREATE TABLE posts (
		id SERIAL PRIMARY KEY,
		author TEXT NOT NULL,
		text TEXT NOT NULL,
		likes INT DEFAULT 0
	);
	CREATE TABLE post_likes (
		post_id INT NOT NULL,
		user_id INT NOT NULL,
		PRIMARY KEY (post_id, user_id)
	);
	`)
	if err != nil {
		log.Fatalf("failed to create schema: %s", err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestPostLifecycle(t *testing.T) {
	// 1. Создаем пост
	post := models.Post{
		Author: "Mishka",
		Text:   "Testoviy test",
	}
	if err := service.CreatePost(post); err != nil {
		t.Fatalf("failed to create post: %v", err)
	}

	// 2. Проверяем список постов
	posts := service.GetPostList()
	if len(posts) != 1 {
		t.Fatalf("expected 1 post, got %d", len(posts))
	}

	// 3. Ставим лайк
	if err := service.LikePost(posts[0].Id, 42); err != nil {
		t.Fatalf("failed to like post: %v", err)
	}

	// 4. Проверяем, что лайк учтен
	updatedPosts := service.GetPostList()
	if updatedPosts[0].Likes != 1 {
		t.Fatalf("expected 1 like, got %d", updatedPosts[0].Likes)
	}
}
