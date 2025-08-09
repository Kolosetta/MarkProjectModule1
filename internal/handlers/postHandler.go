package handlers

import (
	"MarkProjectModule1/internal/models"
	"MarkProjectModule1/internal/service/post"
	responsePkg "MarkProjectModule1/pkg"
	"MarkProjectModule1/pkg/request"
	"net/http"
)

type PostHandler struct{}

// Установка хенжлеров в преданный маршрутиризатор
func RegisterPostHandlers(router *http.ServeMux) {
	handler := PostHandler{}
	router.HandleFunc("POST /posts", handler.CreatePost())
	router.HandleFunc("GET /posts", handler.GetPostList())
}

func (handler *PostHandler) CreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
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
}

func (handler *PostHandler) GetPostList() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		postService := post.NewService(post.GetRepository())
		resultList := postService.GetPostList()
		responsePkg.MakeJsonResponse(w, resultList, http.StatusOK)
	}
}
