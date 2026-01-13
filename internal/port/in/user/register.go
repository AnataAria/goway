package user

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RegisterResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type GetUserRequest struct {
	UserID string
}

type GetUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type RegisterUseCase interface {
	Register(req *RegisterRequest) (*RegisterResponse, error)
}

type GetUserUseCase interface {
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
}
