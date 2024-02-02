package user

import (
	"github.com/EMSI-zero/go-chat/controller/rest/user"
	"github.com/gin-gonic/gin"
)

const (
	LoginPath    = "/login"
	RegisterPath = "/register"
)

func AddRoutes(r *gin.RouterGroup, controller *user.UserController, setAuthPolicyFunc func(string)) {
	r.POST(LoginPath, controller.Login)
	setAuthPolicyFunc(r.BasePath() + LoginPath)
	r.POST(RegisterPath, controller.Register)
	setAuthPolicyFunc(r.BasePath() + RegisterPath)

	userGroup := r.Group("/users")

	userGroup.GET("/", controller.SearchUser)
	userGroup.GET("/:user_id", controller.GetUser)
	userGroup.PATCH("/:user_id", controller.UpdateProfile)
	userGroup.DELETE(":user_id", controller.DeleteAccount)
}
