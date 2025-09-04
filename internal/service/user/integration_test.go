package user

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3/docker"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool
var repo *PostgresRepository
var svc *Service

// Поднимаем PostgreSQL перед запуском тестов
func TestMain(m *testing.M) {
	// Создаем пул докер-ресурсов
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Запускаем контейнер PostgreSQL
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15", // версия
		Env: []string{
			"POSTGRES_USER=tr00per",
			"POSTGRES_PASSWORD=my_pass",
			"POSTGRES_DB=testdb",
		},
	}, func(config *docker.HostConfig) {
		// Автоудаление после остановки
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// Чистим контейнер после тестов
	deferFunc := func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}

	// Формируем dsn
	dsn := fmt.Sprintf("postgres://tr00per:my_pass@localhost:%s/testdb?sslmode=disable", resource.GetPort("5432/tcp"))

	// Ждём, пока PostgreSQL поднимется
	if err := pool.Retry(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		dbPool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}
		return dbPool.Ping(ctx)
	}); err != nil {
		deferFunc()
		log.Fatalf("can't connect to docker: %s", err)
	}

	// Создаем таблицу users
	_, err = dbPool.Exec(context.Background(), `
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL
	)`)
	if err != nil {
		deferFunc()
		log.Fatalf("can't create table: %s", err)
	}

	// Инициализируем репо и сервис
	repo = NewPostgresRepository(dbPool)
	svc = NewService(repo)

	code := m.Run()

	// Чистим контейнер
	deferFunc()

	os.Exit(code)
}

func TestUserLifecycle(t *testing.T) {
	// Создание пользователя
	err := svc.CreateUser("mishka", "blabla@mail.com")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Получчение пользователя
	u, err := svc.GetUser("mishka")
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}
	if u.Username != "mishka" || u.Email != "blabla@mail.com" {
		t.Errorf("unexpected user: %+v", u)
	}

	// Получение всех
	users, err := svc.GetUsersList()
	if err != nil {
		t.Fatalf("failed to get users list: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("expected 1 user, got %d", len(users))
	}
}
