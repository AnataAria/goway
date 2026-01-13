package user

import (
	domainUser "github.com/AnataAria/goway/internal/domain/user"
	portUser "github.com/AnataAria/goway/internal/port/in/user"
)

func (h *UserUseCase) GetUser(req *portUser.GetUserRequest) (*portUser.GetUserResponse, error) {
	user, err := h.userRepo.FindByID(req.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domainUser.NewErrUserNotFound()
	}

	return &portUser.GetUserResponse{
		ID:    user.GetID(),
		Email: user.GetEmail().Value(),
		Name:  user.GetName(),
	}, nil
}
