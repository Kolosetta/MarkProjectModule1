package models

type ResponseBody struct {
	Success bool
}

type RegistrationRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type CreatePostRequest struct {
	Author string `json:"author" validate:"required"`
	Text   string `json:"text" validate:"required"`
}

type LikePostRequest struct {
	UserId int64 `json:"user_id" validate:"required"`
}
