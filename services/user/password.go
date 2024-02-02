package user

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string, salt string) (string, error) {
	passwordWithSalt := password + salt

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func IsValidPassword(password string) (valid bool, err error) {
	uppercaseRegex := "[A-Z]"
	lowercaseRegex := "[a-z]"
	digitRegex := "\\d"
	specialRegex := "[^a-zA-Z0-9]"
	lengthRegex := "^.{8,}$"

	matchUppercase, _ := regexp.MatchString(uppercaseRegex, password)
	if !matchUppercase {
		return false, fmt.Errorf("must have at least one uppercase letter")
	}

	matchLowercase, _ := regexp.MatchString(lowercaseRegex, password)
	if !matchLowercase {
		return false, fmt.Errorf("must have at least one lowercase letter")
	}

	matchDigit, _ := regexp.MatchString(digitRegex, password)
	if !matchDigit {
		return false, fmt.Errorf("must have at least one digit")
	}

	matchSpecial, _ := regexp.MatchString(specialRegex, password)
	if !matchSpecial {
		return false, fmt.Errorf("must have at least one special character")
	}

	matchLength, _ := regexp.MatchString(lengthRegex, password)
	if !matchLength {
		return false, fmt.Errorf("must be at least 8 characters long")
	}

	if err != nil {
		return false, err
	}

	return true, err
}


func validatePassword(inputPassword, storedPassword, salt string) bool {
	inputPasswordWithSalt := inputPassword + salt

	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPasswordWithSalt))
	return err == nil
}

func GenerateSalt(length int) (string, error) {
	byteSize := length / 2

	randomBytes := make([]byte, byteSize)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	salt := hex.EncodeToString(randomBytes)

	return salt, nil
}
