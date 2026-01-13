package user

import (
	"errors"

	"github.com/AnataAria/goway/internal/adapter/in/http/middleware"
	portUser "github.com/AnataAria/goway/internal/port/in/user"
	"github.com/go-fuego/fuego"
)

func (h *UserAdapter) Profile(c fuego.ContextNoBody) (GetUserResponse, error) {
	userID, ok := c.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		return GetUserResponse{}, errors.New("")
	}

	response, err := h.getUserUseCase.GetUser(&portUser.GetUserRequest{
		UserID: userID,
	})
	if err != nil {
		return GetUserResponse{}, err
	}

	return toGetUserResponse(response), nil
}
