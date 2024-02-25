package socket

import (
	"fmt"
	"net/http"
	"time"

	"github.com/EMSI-zero/go-chat/domain/user"
	"github.com/gin-gonic/gin"
)

func HandleTokenAuthentication(c *gin.Context) {
	// Generate a temporary external authentication token
	externalAuthToken, err := generateTemporaryToken()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error generating external authentication token"})
		return
	}

	// Store the external token in the authentication cache
	authCache.Lock()
	authCache.tokens[externalAuthToken.Token] = externalAuthToken.ExpireTime
	authCache.Unlock()

	// Return the external authentication token to the client
	c.String(http.StatusOK, externalAuthToken.Token)
}

func handleWebSocket(hub *Hub, c *gin.Context) {
	// Extract external authentication token from the query parameters
	externalAuthToken := c.Query("token")
	if externalAuthToken == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "External authentication token missing"})
		return
	}

	// Validate the external token against the authentication cache
	authCache.Lock()
	expireTime, found := authCache.tokens[externalAuthToken]
	authCache.Unlock()

	if !found || time.Now().After(expireTime) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid external authentication token"})
		return
	}

	userId, err := user.GetUserFromCtx(ctx)
	if err != nil {
		return
	}


	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}

	
	wsHandler.ServeHTTP(c.Writer, c.Request)
}
