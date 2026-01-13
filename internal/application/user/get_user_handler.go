package user

import (
	domainUser "github.com/AnataAria/goway/internal/domain/user"
	portUser "github.com/AnataAria/goway/internal/port/in/user"
	persistUser "github.com/AnataAria/goway/internal/port/out/persistence/user"
)

type GetUserHandler struct {
	userRepo persistUser.UserRepository
}

func NewGetUserHandler(userRepo persistUser.UserRepository) *GetUserHandler {
	return &GetUserHandler{
		userRepo: userRepo,
	}
}

func (h *GetUserHandler) GetUser(req *portUser.GetUserRequest) (*portUser.GetUserResponse, error) {
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
