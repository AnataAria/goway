package user

import portUser "github.com/AnataAria/goway/internal/port/in/user"

func toGetUserResponse(user *portUser.GetUserResponse) GetUserResponse {
	return GetUserResponse{
		ID:      user.ID,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
		Name:    user.Name,
	}
}

func (r *RegisterRequest) toRegisterInput() *portUser.RegisterRequest {
	return &portUser.RegisterRequest{
		Email:    r.Email,
		Password: r.Password,
		Phone:    r.Phone,
		Address:  r.Address,
		Name:     r.Name,
	}
}

func toRegisterResponse(user *portUser.RegisterResponse) RegisterResponse {
	return RegisterResponse{
		ID:      user.ID,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
		Name:    user.Name,
	}
}
