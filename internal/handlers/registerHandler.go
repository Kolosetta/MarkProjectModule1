package handlers

import (
	"MarkProjectModule1/internal/models"
	"MarkProjectModule1/internal/service/user"
	responsePkg "MarkProjectModule1/pkg"
	"MarkProjectModule1/pkg/request"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type RegHandler struct{}

func RegisterRegHandlers(router *mux.Router) {
	handler := RegHandler{}
	router.HandleFunc("/auth/register", handler.Register).Methods("POST")
}

func (handler *RegHandler) Register(w http.ResponseWriter, req *http.Request) {
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
