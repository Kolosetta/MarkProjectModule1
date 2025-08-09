package handlers

import (
	"MarkProjectModule1/internal/models"
	"MarkProjectModule1/internal/service/user"
	responsePkg "MarkProjectModule1/pkg"
	"MarkProjectModule1/pkg/request"
	"fmt"
	"net/http"
)

type RegHandler struct{}

// Установка хенжлеров в преданный маршрутиризатор
func RegisterRegHandlers(router *http.ServeMux) {
	handler := RegHandler{}
	router.HandleFunc("POST /auth/register", handler.Register())
}

// Метод-обертка, который возвращает функцию-обработчик.
func (handler *RegHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var payload, err = request.HandleBody[models.RegistrationRequest](&w, req)
		if err != nil {
			responsePkg.MakeJsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		userService := user.NewService(user.GetRepository())
		err = userService.CreateUser(payload.Username, payload.Email)

		if err != nil {
			responsePkg.MakeJsonResponse(w, err.Error(), http.StatusBadRequest)
		}

		fmt.Print(payload)
		response := models.ResponseBody{
			Success: true,
		}
		responsePkg.MakeJsonResponse(w, response, http.StatusOK)
	}
}
