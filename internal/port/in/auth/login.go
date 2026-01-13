package auth

type LoginUseCase interface {
	Login(req *LoginRequest) (*LoginResponse, error)
}
