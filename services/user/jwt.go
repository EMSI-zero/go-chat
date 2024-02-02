package user

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret")

type TokenType string

const (
	JWT      TokenType = "JWT"
	REFRESH  TokenType = "Refresh"
	Register TokenType = "Register"
)

var Timeouts = map[TokenType]time.Duration{
	JWT:      time.Second * 60 * 1000,
	REFRESH:  time.Second * 60 * 60 * 24 * 7,
	Register: time.Second * 60 * 60 * 2,
}

func generateToken(userId int64, payload string, tokenType TokenType) (string, error) {
	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(Timeouts[tokenType]).Unix(), // Token expires after timeout
		"iat":     time.Now().Unix(),
		"type":    tokenType,
		"userId":  userId,
		"payload": payload,
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateAndParseToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if (*claims)["exp"].(float64) < float64(time.Now().Unix()) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}

func GetUserIdFromToken(tokenString string) (int64, error) {
	claims, err := validateAndParseToken(tokenString)
	if err != nil {
		return 0, err
	}
	userIDRaw, ok := (*claims)["userId"]
	if !ok {
		return 0, fmt.Errorf("userId claim not found in token")
	}
	userID, ok := userIDRaw.(float64)
	if !ok {
		return 0, fmt.Errorf("userId claim is not a valid float64 in token")
	}
	return int64(userID), nil
}

func GetPayloadFromToken(tokenString string) (string, error) {

	claims, err := validateAndParseToken(tokenString)
	if err != nil {
		return "", err
	}

	payload, ok := (*claims)["payload"].(string)
	if !ok {
		return "", fmt.Errorf("payload claim not found in token")
	}

	return payload, nil
}
