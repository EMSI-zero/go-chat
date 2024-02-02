package user

import (
	"net/http"
	"strconv"

	"github.com/EMSI-zero/go-chat/domain/user"
	"github.com/EMSI-zero/go-chat/infra/httputils"
	"github.com/EMSI-zero/go-chat/registry"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService user.UserService
}

func NewUserController(sr registry.ServiceRegistry) *UserController {
	return &UserController{UserService: sr.GetUserService()}
}

func (uc *UserController) Register(c *gin.Context) {
	ctx := c.Request.Context()

	var req *user.RegisterRequest
	err := c.Bind(req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
		return
	}

	err = uc.UserService.Register(ctx, req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, nil)
}

func (uc *UserController) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req *user.LoginRequest
	err := c.Bind(req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
		return
	}

	res, err := uc.UserService.Login(ctx, req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, res)
}

func (uc *UserController) Logout(c *gin.Context) {
	ctx := c.Request.Context()

	err := uc.UserService.Logout(ctx)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, nil)
}

func (uc *UserController) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	res, err := uc.UserService.GetUser(ctx, int64(userID))
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, res)
}

func (uc *UserController) UpdateProfile(c *gin.Context) {
	ctx := c.Request.Context()

	var req *user.UpdateUserRequest
	err := c.Bind(req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
		return
	}

	err = uc.UserService.UpdateUser(ctx, req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, nil)
}

func (uc *UserController) DeleteAccount(c *gin.Context) {
	ctx := c.Request.Context()

	var req *user.DeleteUserRequest
	err := c.Bind(req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
		return
	}

	err = uc.UserService.DeleteUser(ctx, req)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, nil)
}

func (uc *UserController) SearchUser(c *gin.Context) {
	ctx := c.Request.Context()

	keyword := c.Param("keyword")

	res, err := uc.UserService.GetUserByKeyWord(ctx, keyword)
	if err != nil {
		httputils.NewError(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, res)
}
