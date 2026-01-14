package user

import (
	"errors"

	"github.com/AnataAria/goway/internal/adapter/in/http/middleware"
	portUser "github.com/AnataAria/goway/internal/port/in/user"
	"github.com/AnataAria/goway/pkg/response"
	"github.com/go-fuego/fuego"
)

func (h *UserAdapter) Profile(c fuego.ContextNoBody) (*response.Response[GetUserResponse], error) {
	userID, ok := c.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, errors.New("user is not authenticated")
	}

	res, err := h.userUseCase.GetUser(&portUser.GetUserRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	return response.RespondOK(toGetUserResponse(res)), nil
}
