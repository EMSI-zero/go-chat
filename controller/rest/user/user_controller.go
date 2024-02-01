package user

import (
	"net/http"

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
