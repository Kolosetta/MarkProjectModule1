package main

import (
	"MarkProjectModule1/internal/handlers"
	"MarkProjectModule1/internal/service/post"
	"MarkProjectModule1/pkg/events"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()            //Создаем маршрутиризатор для сервера
	handlers.RegisterRegHandlers(router) //Регистрируем в маршрутиризаторе хендлер дял ендпоинта /auth/register
	handlers.RegisterPostHandlers(router)

	//Запускаем нвоый воркер, который читает очередь лайков
	postService := post.NewService(post.GetRepository())
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
