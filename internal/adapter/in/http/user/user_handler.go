package user

import (
	portUser "github.com/AnataAria/goway/internal/port/in/user"
	"github.com/AnataAria/goway/pkg/response"
	"github.com/go-fuego/fuego"
)

func (h *UserAdapter) Register(c fuego.ContextWithBody[RegisterRequest]) (*response.Response[RegisterResponse], error) {
	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	res, err := h.userUseCase.Register(body.toRegisterInput())
	if err != nil {
		return nil, err
	}

	return response.RespondCreated(toRegisterResponse(res)), nil
}

func (h *UserAdapter) GetUser(c fuego.ContextNoBody) (*response.Response[GetUserResponse], error) {
	userID := c.PathParam("id")

	res, err := h.userUseCase.GetUser(&portUser.GetUserRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	return response.RespondOK(toGetUserResponse(res)), nil
}

func (h *UserAdapter) GetUserByID(c fuego.ContextNoBody) (*response.Response[GetUserResponse], error) {
	userID := c.PathParam("id")

	res, err := h.userUseCase.GetUser(&portUser.GetUserRequest{
		UserID: userID,
	})

	if err != nil {
		return nil, err
	}

	return response.RespondOK(toGetUserResponse(res)), nil
}
