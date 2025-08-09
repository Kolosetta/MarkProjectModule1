package main

import (
	"MarkProjectModule1/internal/handlers"
	"net/http"
)

func main() {

	router := http.NewServeMux()         //Создаем маршрутиризатор для сервера
	handlers.RegisterRegHandlers(router) //Регистрируем в маршрутиризаторе хендлер дял ендпоинта /auth/register
	handlers.RegisterPostHandlers(router)

	//конфигурируем сервер. Назначаем роутер, котоырй будет распределять запросы
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	// Запускаем сервер
	server.ListenAndServe()

}
