package user

import (
	"github.com/EMSI-zero/go-chat/controller/rest"
	"github.com/EMSI-zero/go-chat/gateway/user/contact"
	"github.com/gin-gonic/gin"
)

const (
	LoginPath    = "/login"
	RegisterPath = "/register"
)

func AddRoutes(r *gin.RouterGroup, controller *rest.Controller, setAuthPolicyFunc func(string)) {
	r.POST(LoginPath, controller.UserController.Login)
	setAuthPolicyFunc(r.BasePath() + LoginPath)
	r.POST(RegisterPath, controller.UserController.Register)
	setAuthPolicyFunc(r.BasePath() + RegisterPath)

	userGroup := r.Group("/users")

	userGroup.GET("/", controller.UserController.SearchUser)
	userGroup.GET("/:user_id", controller.UserController.GetUser)
	userGroup.PATCH("/:user_id", controller.UserController.UpdateProfile)
	userGroup.DELETE(":user_id", controller.UserController.DeleteAccount)

	contact.AddRoutes(userGroup, controller, setAuthPolicyFunc)
}
