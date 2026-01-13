package user

import (
	domainUser "github.com/AnataAria/goway/internal/domain/user"
	portUser "github.com/AnataAria/goway/internal/port/in/user"
	persistUser "github.com/AnataAria/goway/internal/port/out/persistence/user"
	"github.com/google/uuid"
)

type RegisterHandler struct {
	userRepo persistUser.UserRepository
}

func NewRegisterHandler(userRepo persistUser.UserRepository) *RegisterHandler {
	return &RegisterHandler{
		userRepo: userRepo,
	}
}

func (h *RegisterHandler) Register(req *portUser.RegisterRequest) (*portUser.RegisterResponse, error) {
	email, err := domainUser.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}

	password, err := domainUser.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}

	exists, err := h.userRepo.ExistsByEmail(email.Value())
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domainUser.NewErrUserAlreadyExists()
	}

	userID := uuid.New().String()
	newUser := domainUser.NewUser(userID, email, password, req.Name)

	err = h.userRepo.Save(newUser)
	if err != nil {
		return nil, err
	}

	return &portUser.RegisterResponse{
		ID:    newUser.GetID(),
		Email: newUser.GetEmail().Value(),
		Name:  newUser.GetName(),
	}, nil
}
