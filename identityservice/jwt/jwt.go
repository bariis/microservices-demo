package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"

	"github.com/google/uuid"
)

type token struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

type TokenResponse struct {
	AccessToken token `json:"access_token"`
}

func GenerateToken(userID uuid.UUID) token {
	return token{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Audience:  "http://krakend:5000",
			Issuer:    "http://identityservice:5002",
		},
	}
}
