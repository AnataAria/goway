package user

import (
	portUser "github.com/AnataAria/goway/internal/port/in/user"
	"github.com/go-fuego/fuego"
)

func (h *UserAdapter) Register(c fuego.ContextWithBody[RegisterRequest]) (RegisterResponse, error) {
	body, err := c.Body()
	if err != nil {
		return RegisterResponse{}, err
	}

	response, err := h.registerUseCase.Register(body.toRegisterInput())
	if err != nil {
		return RegisterResponse{}, err
	}

	return toRegisterResponse(response), nil
}

func (h *UserAdapter) GetUser(c fuego.ContextNoBody) (GetUserResponse, error) {
	userID := c.PathParam("id")

	response, err := h.getUserUseCase.GetUser(&portUser.GetUserRequest{
		UserID: userID,
	})
	if err != nil {
		return GetUserResponse{}, err
	}

	return toGetUserResponse(response), nil
}

func (h *UserAdapter) GetUserByID(c fuego.ContextNoBody) (GetUserResponse, error) {
	userID := c.PathParam("id")

	response, err := h.getUserUseCase.GetUser(&portUser.GetUserRequest{
		UserID: userID,
	})

	if err != nil {
		return GetUserResponse{}, err
	}

	return toGetUserResponse(response), nil
}
