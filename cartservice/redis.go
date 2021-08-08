package main

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"strings"
)

// gets and returns the required part of the token
func getUserIdFromToken(token string) (string, error) {
	// remove the `bearer` part from the token
	encodedToken := strings.Split(token, " ")[1]
	parsedToken, _ := jwt.Parse(encodedToken, nil)
	if parsedToken == nil {
		return "", errors.New("token is not valid")
	}

	claims, _ := parsedToken.Claims.(jwt.MapClaims)

	return claims["user_id"].(string), nil
}
