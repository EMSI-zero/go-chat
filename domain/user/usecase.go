package user

import (
	"context"
	"fmt"
)

type UserService interface {
	mustImplementBaseService()
	Register(ctx context.Context, req *RegisterRequest) error
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	Logout(ctx context.Context) error
	GetCurrentUser(ctx context.Context) (*GetCurrentUserResponse, error)
	GetUser(ctx context.Context, id int64) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest) error
	DeleteUser(ctx context.Context, req *DeleteUserRequest) error
	GetUserByKeyWord(ctx context.Context, keyword string) ([]*GetUserResponse, error)
	ValidateToken(context.Context, string) (int64, error)
	RefreshToken(context.Context, string) (string, error)
}

type UnImplementedUserService struct{}

func (*UnImplementedUserService) mustImplementBaseService() {}

func (*UnImplementedUserService) Register(ctx context.Context, req RegisterRequest) error {
	return fmt.Errorf("service not implemented")
}

func (*UnImplementedUserService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, fmt.Errorf("service not implemented")
}

func (*UnImplementedUserService) Logout(ctx context.Context) error {
	return fmt.Errorf("service not implemented")
}

func (*UnImplementedUserService) GetCurrentUser(ctx context.Context) (*GetCurrentUserResponse, error) {
	return nil, fmt.Errorf("service not implemented")
}

func (*UnImplementedUserService) GetUser(ctx context.Context, id int64) (*GetUserResponse, error) {
	return nil, fmt.Errorf("service not implemented")
}

func (*UnImplementedUserService) UpdateUser(ctx context.Context, req *UpdateUserRequest) error {
	return fmt.Errorf("service not implemented")
}

func (*UnImplementedUserService) DeleteUser(ctx context.Context, req *DeleteUserRequest) error {
	return fmt.Errorf("service not implemented")
}

func (*UnImplementedUserService) GetUserByKeyWord(ctx context.Context, keyword string) ([]*GetUserResponse, error) {
	return nil, fmt.Errorf("service not implemented")
}

func (*UnImplementedUserService) ValidateToken(context.Context, string) (int64, error) {
	return 0, fmt.Errorf("service not implemented")
}

func (*UnImplementedUserService) RefreshToken(context.Context, string) (string, error) {
	return "", fmt.Errorf("service not implemented")
}
