package main

import (
	"MarkProjectModule1/internal/handlers"
	"MarkProjectModule1/internal/service/post"
	"MarkProjectModule1/internal/service/user"
	"MarkProjectModule1/pkg/events"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	//Запуск профилирования на отдельном порту
	go func() {
		log.Println("pprof запущен на :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	router := mux.NewRouter()                                                    //Создаем маршрутиризатор для сервера
	handlers.RegisterRegHandlers(router, user.NewService(user.GetRepository()))  //Регистрируем в маршрутиризаторе хендлер дял ендпоинта /auth/register
	handlers.RegisterPostHandlers(router, post.NewService(post.GetRepository())) // передаем внутрь инстансы сервисов

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
