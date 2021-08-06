package main

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

type product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}

func parsePayload(token string) (string, error) {
	parsedToken, _ := jwt.Parse(token, nil)
	if parsedToken == nil {
		return "", errors.New("error: empty token")
	}

	claims, _ := parsedToken.Claims.(jwt.MapClaims)

	return claims["user_id"].(string), nil
}
