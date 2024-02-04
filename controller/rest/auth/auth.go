package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/EMSI-zero/go-chat/domain/user"
	"github.com/EMSI-zero/go-chat/registry"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService user.UserService
	authByPass  map[string]bool
}

func NewAuthController(sr registry.ServiceRegistry) *AuthController {
	return &AuthController{
		userService: sr.GetUserService(),
		authByPass: make(map[string]bool),
	}
}

func (ac *AuthController) SetByPassPolicy(route string) {
	ac.authByPass[route] = true
}

func (ac *AuthController) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if ac.authByPass[c.FullPath()] {
			c.Next()
			return
		}

		ctx := c.Request.Context()

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or missing Bearer Token"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		userId, err := ac.userService.ValidateToken(ctx, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("no user id found"))
			log.Print(err)
			return
		}

		ctx = context.WithValue(ctx, user.UserContextKey{}, userId)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
