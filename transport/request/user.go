package request

type RegisterUserRequest struct {
	Name       string `json:"name" validate:"required"`
	Occupation string `json:"occupation" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8"`
}
