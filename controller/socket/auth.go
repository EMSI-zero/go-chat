package socket

import (
	"sync"
	"time"
)

type AuthenticationCache struct {
	sync.Mutex
	tokens map[string]time.Time
}

var authCache *AuthenticationCache

type Ticket struct {
	Token      string
	ExpireTime time.Time
}

func InitAuthenticationCache() {
	authCache = &AuthenticationCache{
		tokens: make(map[string]time.Time),
	}
}

func generateRandomToken() string {
	// Implement your logic to generate a secure random token
	return "randomtoken123"
}

func generateTemporaryToken() (*Ticket, error) {
	token := generateRandomToken()
	expireTime := time.Now().Add(time.Minute * 5)
	return &Ticket{Token: token, ExpireTime: expireTime}, nil
}

func validateTemporaryToken(ticket *Ticket) bool {
	return time.Now().Before(ticket.ExpireTime)
}
