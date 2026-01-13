package user

type RegisterRequest struct {
	Email    string
	Password string
	Phone    string
	Address  string
	Name     string
}

type RegisterResponse struct {
	ID      string
	Email   string
	Phone   string
	Address string
	Name    string
}

type GetUserRequest struct {
	UserID string
}

type GetUserResponse struct {
	ID      string
	Email   string
	Phone   string
	Address string
	Name    string
}
