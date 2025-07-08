package validators

type CreateAdminValidator struct {
	Username string `json:"username" validate:"required,min=4,max=50"`
	Password string `json:"password" validate:"required,min=6"`
	Phone    string `json:"phone" validate:"required,e164"`
}