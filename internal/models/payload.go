package models

type RegisterRes struct {
	Success bool
}

type RegistrationRequest struct {
	Name    string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}
