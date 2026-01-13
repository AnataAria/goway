package user

type UserUseCase interface {
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
	Register(req *RegisterRequest) (*RegisterResponse, error)
}
