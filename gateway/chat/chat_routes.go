package chat

import (
	"github.com/EMSI-zero/go-chat/controller/rest"
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.RouterGroup, controller *rest.Controller, setAuthPolicyFunc func(string)) {
	cRoutes := r.Group("/chats")

	cRoutes.GET("/", controller.ChatController.GetChats)
	cRoutes.GET("/:chat_id", controller.ChatController.GetChat)
	cRoutes.POST("/", controller.ChatController.NewChat)
	cRoutes.DELETE("/:chat_id", controller.ChatController.LeaveChat)
	cRoutes.DELETE("/:chat_id/messages/:message_id", controller.ChatController.DeleteMessage)
}
