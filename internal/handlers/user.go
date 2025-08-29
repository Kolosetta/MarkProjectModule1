package handlers

import (
	"MarkProjectModule1/internal/models"
	"MarkProjectModule1/internal/service/user"
	responsePkg "MarkProjectModule1/pkg"
	"MarkProjectModule1/pkg/request"
	"github.com/gorilla/mux"
	"net/http"
)

type RegHandler struct {
	service *user.Service
}

func RegisterRegHandlers(router *mux.Router, service *user.Service) {
	handler := RegHandler{service: service}
	router.HandleFunc("/auth/register", handler.Register).Methods("POST")
	router.HandleFunc("/user/{username}", handler.GetUser).Methods("GET")
	router.HandleFunc("/users", handler.GetUsersList).Methods("GET")
}

func (handler *RegHandler) Register(w http.ResponseWriter, req *http.Request) {
	var payload, err = request.HandleBody[models.RegistrationRequest](&w, req)
	if err != nil {
		responsePkg.MakeJsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.service.CreateUser(payload.Username, payload.Email)

	if err != nil {
		responsePkg.MakeJsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := models.ResponseBody{
		Success: true,
	}
	responsePkg.MakeJsonResponse(w, response, http.StatusOK)
}

func (handler *RegHandler) GetUser(w http.ResponseWriter, req *http.Request) {
	var queryVars = mux.Vars(req)
	username := queryVars["username"]
	userObj, err := handler.service.GetUser(username)

	if err != nil {
		responsePkg.MakeJsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responsePkg.MakeJsonResponse(w, userObj, http.StatusOK)

}

func (handler *RegHandler) GetUsersList(w http.ResponseWriter, req *http.Request) {
	users, err := handler.service.GetUsersList()
	if err != nil {
		responsePkg.MakeJsonResponse(w, err.Error(), http.StatusBadRequest)
	}
	responsePkg.MakeJsonResponse(w, users, http.StatusOK)
}
