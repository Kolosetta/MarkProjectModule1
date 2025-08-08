package models

type RegisterRes struct {
	Success bool
}

type RegistrationRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}
