package user

import (
	"context"

	"github.com/EMSI-zero/go-chat/domain/user"
	"github.com/EMSI-zero/go-chat/infra/dbrepo"
	"github.com/EMSI-zero/go-chat/registry"
)

type UserService struct {
	user.UnImplementedUserService
	db dbrepo.DBConn
}

func NewUserService(registry registry.ServiceRegistry) *UserService {
	return &UserService{
		db: registry.GetDB(),
	}
}

func (us *UserService) Register(ctx context.Context, req user.RegisterRequest) error {
	db, err := us.db.GetConn(ctx)
	if err != nil {
		return nil
	}

	user := &user.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Password:  req.Password,
		Bio:       req.Bio,
		ImagePath: req.ImagePath,
	}

	err = db.Create(user).Error
	if err != nil {
		return err
	}
	
	return nil
}
func (*UserService) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	return nil, nil
}
func (*UserService) Logout(ctx context.Context) error {
	return nil
}
func (*UserService) GetCurrentUser(ctx context.Context) (*user.GetCurrentUserResponse, error) {
	return nil, nil
}
func (*UserService) GetUser(ctx context.Context, id int64) (*user.GetUserResponse, error) {
	return nil, nil
}
func (*UserService) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) error {
	return nil
}
func (*UserService) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) {
	return
}
