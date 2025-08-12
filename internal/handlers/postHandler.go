package handlers

import (
	"MarkProjectModule1/internal/models"
	"MarkProjectModule1/internal/service/post"
	responsePkg "MarkProjectModule1/pkg"
	"MarkProjectModule1/pkg/request"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PostHandler struct {
	service *post.Service
}

// Установка хендлеров в преданный маршрутиризатор
func RegisterPostHandlers(router *mux.Router, service *post.Service) {
	handler := PostHandler{service: service}
	router.HandleFunc("/posts", handler.CreatePost).Methods("POST")
	router.HandleFunc("/posts", handler.GetPostList).Methods("GET")
	router.HandleFunc("/posts/{id}/like", handler.LikePost).Methods("POST")
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, req *http.Request) {
	var payload, err = request.HandleBody[models.CreatePostRequest](&w, req)
	if err != nil {
		responsePkg.MakeJsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.service.CreatePost(models.Post{
		Author: payload.Author,
		Text:   payload.Text,
	})
	if err != nil {
		responsePkg.MakeJsonResponse(w, err.Error(), http.StatusBadRequest)
	}

	response := models.ResponseBody{Success: true}
	responsePkg.MakeJsonResponse(w, response, http.StatusOK)
}

func (handler *PostHandler) GetPostList(w http.ResponseWriter, req *http.Request) {
	resultList := handler.service.GetPostList()
	responsePkg.MakeJsonResponse(w, resultList, http.StatusOK)
}

func (handler *PostHandler) LikePost(w http.ResponseWriter, req *http.Request) {
	var payload, err = request.HandleBody[models.LikePostRequest](&w, req)
	if err != nil {
		responsePkg.MakeJsonResponse(w, err.Error(), http.StatusBadRequest)
	}
	queryVars := mux.Vars(req)
	postId, err := strconv.ParseInt(queryVars["id"], 10, 64)

	err = post.AddLikeAsync(postId, payload.UserId)

	//Если очередь переполнена
	if err != nil {
		responsePkg.MakeJsonResponse(w, err.Error(), http.StatusTooManyRequests)
	}
}
