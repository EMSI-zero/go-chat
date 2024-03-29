package rest

import (
	"log"

	"github.com/EMSI-zero/go-chat/controller/rest/auth"
	"github.com/EMSI-zero/go-chat/controller/rest/chat"
	"github.com/EMSI-zero/go-chat/controller/rest/contact"
	"github.com/EMSI-zero/go-chat/controller/rest/user"
	"github.com/EMSI-zero/go-chat/registry"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	AuthController    *auth.AuthController
	UserController    *user.UserController
	ContactController *contact.ContactController
	ChatController    *chat.ChatController
}

func NewController(sr registry.ServiceRegistry) *Controller {
	controller := new(Controller)

	controller.AuthController = auth.NewAuthController(sr)
	controller.UserController = user.NewUserController(sr)
	controller.ContactController = contact.NewContactController(sr)
	controller.ChatController = chat.NewChatController(sr)

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
