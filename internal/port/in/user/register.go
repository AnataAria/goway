package user

type RegisterUseCase interface {
	Register(req *RegisterRequest) (*RegisterResponse, error)
}

type GetUserUseCase interface {
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
}
