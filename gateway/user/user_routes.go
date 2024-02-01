package user

import "github.com/gin-gonic/gin"

func AddRoutes(r *gin.RouterGroup){
	r.POST("/register")
	r.POST("/login")

	userGroup:= r.Group("/users")

	userGroup.GET("/:user_id")
}