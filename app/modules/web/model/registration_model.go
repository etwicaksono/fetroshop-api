package model

type RegistrationRequest struct {
	Username string `json:"username" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required"`
	FullName string `json:"fullname" validate:"required"`
}