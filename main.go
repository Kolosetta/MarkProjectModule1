package main

import (
	"MarkProjectModule1/internal/handlers"
	"MarkProjectModule1/internal/service/post"
	"MarkProjectModule1/internal/service/user"
	"MarkProjectModule1/pkg/db"
	"MarkProjectModule1/pkg/events"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	//Запуск профилирования на отдельном порту
	events.StartLogger()

	router := mux.NewRouter() //Создаем маршрутиризатор для сервера

	pool, err := db.NewPostgresPool()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	//Регистрируем в маршрутиризаторе хендлер дял ендпоинта /auth/register
	var postService = post.NewService(post.NewPostgresRepository(pool))
	handlers.RegisterPostHandlers(router, postService) // передаем внутрь инстансы сервисов

	var userService = user.NewService(user.NewPostgresRepository(pool))
	handlers.RegisterRegHandlers(router, userService)

	//Запускаем новый воркер, который читает очередь лайков
	post.StartLikeWorker(postService)

	events.StartLogger()

	//конфигурируем сервер. Назначаем роутер, котоырй будет распределять запросы
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	// Запускаем сервер
	server.ListenAndServe()

}
