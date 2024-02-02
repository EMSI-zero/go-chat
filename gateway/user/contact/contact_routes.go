package contact

import (
	"github.com/EMSI-zero/go-chat/controller/rest"
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.RouterGroup, controller *rest.Controller, setAuthPolicyFunc func(string)) {
	cGroup := r.Group("/:user_id/contacts")
	cGroup.POST("/contact", controller.ContactController.AddContact)
	cGroup.GET("/contact", controller.ContactController.GetContacts)
	cGroup.DELETE("/contact/:id", controller.ContactController.RemoveContact)
}
