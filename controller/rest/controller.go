package rest

import (
	"log"

	"github.com/EMSI-zero/go-chat/controller/rest/auth"
	"github.com/EMSI-zero/go-chat/controller/rest/user"
	"github.com/EMSI-zero/go-chat/registry"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	AuthController *auth.AuthController
	UserController *user.UserController
}

func NewController(sr registry.ServiceRegistry) *Controller {
	controller := new(Controller)

	controller.AuthController = auth.NewAuthController(sr)
	controller.UserController = user.NewUserController(sr)

	return controller
}

func (c *Controller) HandleError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) > 0 {
			log.Print(ctx.Errors[0].Error())
			ctx.JSON(400, gin.H{
				"message": ctx.Errors[0].Error(),
			})
		}
	}
}
