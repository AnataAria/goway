package user

import (
	"github.com/go-fuego/fuego"
)

func (h *UserHandler) Profile(c fuego.ContextNoBody) (GetUserResponse, error) {
	userID := c.Context().Value("user_id").(string)

	user, err := h.userRepo.FindByID(userID)
	if err != nil {
		return GetUserResponse{}, err
	}
	if user == nil {
		return GetUserResponse{}, err
	}

	return GetUserResponse{
		ID:    user.GetID(),
		Email: user.GetEmail().Value(),
		Name:  user.GetName(),
	}, nil
}
