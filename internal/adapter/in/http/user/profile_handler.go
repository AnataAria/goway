package user

import (
	portUser "github.com/AnataAria/goway/internal/port/in/user"
	"github.com/go-fuego/fuego"
)

func (h *UserAdapter) Profile(c fuego.ContextNoBody) (GetUserResponse, error) {
	userID := c.Context().Value("user_id").(string)
	response, err := h.getUserUseCase.GetUser(&portUser.GetUserRequest{
		UserID: userID,
	})
	if err != nil {
		return GetUserResponse{}, err
	}

	return toGetUserResponse(response), nil
}
