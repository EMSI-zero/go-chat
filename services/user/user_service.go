package user

import (
	"context"
	"fmt"
	"log"

	"github.com/EMSI-zero/go-chat/domain/user"
	"github.com/EMSI-zero/go-chat/infra/dbrepo"
	"github.com/EMSI-zero/go-chat/registry"
)

type UserService struct {
	user.UserService
	db dbrepo.DBConn
}

func NewUserService(registry registry.ServiceRegistry) {
	registry.RegisterUserService(&UserService{
		db: registry.GetDB(),
	})
}

func (us *UserService) Register(ctx context.Context, req *user.RegisterRequest) error {
	db, err := us.db.GetConn(ctx)
	if err != nil {
		return nil
	}

	salt, err := GenerateSalt(16)
	if err != nil {
		return err
	}

	log.Print(req)
	hashedPassword, err := hashPassword(req.Password, salt)
	if err != nil {
		return err
	}

	user := &user.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		UserName:  req.UserName,
		Phone:     req.Phone,
		Password:  hashedPassword,
		Salt:      salt,
		Bio:       req.Bio,
		ImagePath: req.ImagePath,
	}

	err = db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	db, err := us.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}

	var existingUser user.User
	if err := db.Where("user_name = ?", req.UserName).First(&existingUser).Error; err != nil {
		return nil, fmt.Errorf("login failed: %v", err)
	}

	if !validatePassword(req.Password, existingUser.Password, existingUser.Salt) {
		return nil, fmt.Errorf("login failed: invalid password")
	}

	// Generate JWT token
	token, err := generateToken(existingUser.ID, existingUser.UserName, JWT)
	if err != nil {
		return nil, err
	}

	refresh, err := generateToken(existingUser.ID, existingUser.UserName, REFRESH)
	if err != nil {
		return nil, err
	}

	response := &user.LoginResponse{
		Token:   token,
		Refresh: refresh,
	}

	return response, nil
}

func (*UserService) Logout(ctx context.Context) error {
	return fmt.Errorf("service not implemented")
}

func (us *UserService) GetCurrentUser(ctx context.Context) (*user.GetCurrentUserResponse, error) {
	db, err := us.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}

	uid, err := user.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	currUser := new(user.User)
	err = db.Model(currUser).Find(&currUser, uid).Error
	if err != nil {
		return nil, err
	}

	return &user.GetCurrentUserResponse{
		FirstName: currUser.FirstName,
		LastName:  currUser.LastName,
		UserName:  currUser.UserName,
		Phone:     currUser.Phone,
		Bio:       currUser.Bio,
		ImagePath: currUser.ImagePath,
	}, nil

}

func (us *UserService) GetUser(ctx context.Context, id int64) (*user.GetUserResponse, error) {
	db, err := us.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}

	currUser := new(user.User)
	err = db.Model(currUser).Find(&currUser, id).Error
	if err != nil {
		return nil, err
	}

	return &user.GetUserResponse{
		FirstName: currUser.FirstName,
		LastName:  currUser.LastName,
		UserName:  currUser.UserName,
		Phone:     currUser.Phone,
		Bio:       currUser.Bio,
		ImagePath: currUser.ImagePath,
	}, nil
}

func (us *UserService) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) error {
	db, err := us.db.GetConn(ctx)
	if err != nil {
		return nil
	}

	user := new(user.User)
	err = db.Model(user).Where("user_name = ?", req.UserName).Take(&user).Error
	if err != nil {
		return err
	}

	log.Print(req)
	hashedPassword, err := hashPassword(req.Password, user.Salt)
	if err != nil {
		return err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Phone = req.Phone
	user.Password = hashedPassword
	user.Bio = req.Bio
	user.ImagePath = req.ImagePath

	err = db.Save(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) error {
	db, err := us.db.GetConn(ctx)
	if err != nil {
		return err
	}

	uid, err := user.GetUserFromCtx(ctx)
	if err != nil {
		return err
	}

	if uid != req.ID {
		return fmt.Errorf("unauthorized")
	}

	err = db.Delete(&user.User{}, uid).Error
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUserByKeyWord(ctx context.Context, keyword string) ([]*user.GetUserResponse, error) {
	db, err := us.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*user.User, 0)
	err = db.Model(&user.User{}).Where("user_name LIKE ?%", keyword).Find(&users).Error
	if err != nil {
		return nil, err
	}

	alikeUsers := make([]*user.GetUserResponse, 0)
	for _, u := range users {
		alikeUsers = append(alikeUsers, &user.GetUserResponse{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			UserName:  u.UserName,
			Phone:     u.Phone,
			Bio:       u.Bio,
			ImagePath: u.ImagePath,
		})
	}

	return alikeUsers, nil
}

func (*UserService) ValidateToken(ctx context.Context, token string) (int64, error) {
	userID, err := GetUserIdFromToken(token)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (*UserService) RefreshToken(ctx context.Context, token string) (string, error) {
	tokenClaims, err := validateAndParseToken(token)
	if err != nil {
		return "", err
	}

	userIDRaw, ok := (*tokenClaims)["userId"]
	if !ok {
		return "", fmt.Errorf("userId claim not found in token")
	}
	userID, ok := userIDRaw.(int64)
	if !ok {
		return "", fmt.Errorf("userId claim is not a valid float64 in token")
	}

	newToken, err := generateToken(userID, "refresh", JWT)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return newToken, nil
}
