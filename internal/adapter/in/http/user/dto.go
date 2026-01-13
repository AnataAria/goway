package user

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type RegisterResponse struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Name    string `json:"name"`
}

type GetUserResponse struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Name    string `json:"name"`
}
