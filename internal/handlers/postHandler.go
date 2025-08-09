package handlers

import (
	"MarkProjectModule1/internal/models"
	"MarkProjectModule1/internal/service/post"
	responsePkg "MarkProjectModule1/pkg"
	"MarkProjectModule1/pkg/request"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PostHandler struct{}

// Установка хендлеров в преданный маршрутиризатор
func RegisterPostHandlers(router *mux.Router) {
	handler := PostHandler{}
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

	postService := post.NewService(post.GetRepository())
	err = postService.CreatePost(models.Post{
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
	postService := post.NewService(post.GetRepository())
	resultList := postService.GetPostList()
	responsePkg.MakeJsonResponse(w, resultList, http.StatusOK)
}

func (handler *PostHandler) LikePost(w http.ResponseWriter, req *http.Request) {
	var payload, err = request.HandleBody[models.LikePostRequest](&w, req)
	if err != nil {
		responsePkg.MakeJsonResponse(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Print(payload)
	vars := mux.Vars(req)
	postId, err := strconv.ParseInt(vars["id"], 10, 64)

	postService := post.NewService(post.GetRepository())
	err = postService.LikePost(postId, payload.UserId)
	if err != nil {
		responsePkg.MakeJsonResponse(w, err.Error(), http.StatusBadRequest)
	}
}
