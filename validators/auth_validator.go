package validators

import (
	"github.com/go-playground/validator/v10"
)


var Validate = validator.New()

type RegisterValidator struct{
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Phone    string `json:"phone" validate:"required,e164"` 
}

type LoginValidator struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type ResendVerificationValidator struct {
	Email string `json:"email" validate:"required,email"`
}
