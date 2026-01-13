package user

import (
	"github.com/AnataAria/goway/internal/application/user"
	portUser "github.com/AnataAria/goway/internal/port/in/user"
	persistUser "github.com/AnataAria/goway/internal/port/out/persistence/user"
	"github.com/go-fuego/fuego"
)

type UserHandler struct {
	userRepo        persistUser.UserRepository
	registerUseCase *user.RegisterHandler
	getUserUseCase  *user.GetUserHandler
}

func NewUserHandler(userRepo persistUser.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo:        userRepo,
		registerUseCase: user.NewRegisterUseCase(userRepo),
		getUserUseCase:  user.NewGetUserUseCase(userRepo),
	}
}

func (h *UserHandler) Register(c fuego.ContextWithBody[RegisterRequest]) (RegisterResponse, error) {
	body, err := c.Body()
	if err != nil {
		return RegisterResponse{}, err
	}

	response, err := h.registerUseCase.Register(&portUser.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
		Name:     body.Name,
	})
	if err != nil {
		return RegisterResponse{}, err
	}

	return RegisterResponse{
		ID:    response.ID,
		Email: response.Email,
		Name:  response.Name,
	}, nil
}

func (h *UserHandler) GetUser(c fuego.ContextNoBody) (GetUserResponse, error) {
	userID := "placeholder-user-id"

	response, err := h.getUserUseCase.GetUser(&portUser.GetUserRequest{
		UserID: userID,
	})
	if err != nil {
		return GetUserResponse{}, err
	}

	return GetUserResponse{
		ID:    response.ID,
		Email: response.Email,
		Name:  response.Name,
	}, nil
}

func (h *UserHandler) GetUserByID(c fuego.ContextNoBody) (GetUserResponse, error) {
	userID := c.PathParam("id")

	response, err := h.getUserUseCase.GetUser(&portUser.GetUserRequest{
		UserID: userID,
	})

	if err != nil {
		return GetUserResponse{}, err
	}

	return GetUserResponse{
		ID:    response.ID,
		Email: response.Email,
		Name:  response.Name,
	}, nil
}
