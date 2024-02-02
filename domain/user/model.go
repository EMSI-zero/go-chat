package user

import (
	"context"
	"fmt"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Phone     string `json:"phone"`
	Password  string `json:"password" gorm:"col:password_c"`
	Salt      string `json:"-"`
	Bio       string `json:"bio"`
	ImagePath string `json:"imagePath"`
}

func (User) TableName() string {
	return "chat_user"
}


type UserSummary struct{
	ID        int64  `json:"id"`
	UserName  string `json:"username"`
	ImagePath string `json:"imagePath"`
}

type RegisterRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	ImagePath string `json:"imagePath"`
}

type LoginRequest struct {
	UserName string
	Password string
}

type LoginResponse struct {
	Token   string
	Refresh string
}

type GetCurrentUserResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Phone     string `json:"phone"`
	Bio       string `json:"bio"`
	ImagePath string `json:"imagePath"`
}

type GetUserResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Phone     string `json:"phone"`
	Bio       string `json:"bio"`
	ImagePath string `json:"imagePath"`
}

type UpdateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	ImagePath string `json:"imagePath"`
}

type UpdateUserResponse struct{}

type DeleteUserRequest struct {
	ID int64 `json:"id"`
}

type UserContextKey struct{}

func GetUserFromCtx(ctx context.Context) (int64, error) {
	userIdValue := ctx.Value(UserContextKey{})
	if userIdValue == nil {
		return 0, fmt.Errorf("no user id found")
	}

	return userIdValue.(int64), nil
}
