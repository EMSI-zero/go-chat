package user

import (
	"github.com/EMSI-zero/go-chat/domain/user"
	"github.com/EMSI-zero/go-chat/registry"
)

type UserService struct {
	user.UnImplementedUserService
}

func NewUserService(registry registry.ServiceRegistry) *UserService {
	return &UserService{}
}
