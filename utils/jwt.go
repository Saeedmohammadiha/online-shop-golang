package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token interface {
	NewAccessToken(userID int) (string, error)
	NewRefreshToken(userID int) (string, error)
}

type Jwt struct {
}

func NewAuth() Token {
	return &Jwt{}
}

func (*Jwt) NewAccessToken(userID int) (string, error) {
	var claims = jwt.MapClaims{
		"userID":    userID,
		"IssuedAt":  time.Now().Unix(),
		"ExpiresAt": time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	AccessToken, err := token.SignedString([]byte(GetEnv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return AccessToken, nil
}

func (*Jwt) NewRefreshToken(userID int) (string, error) {
	var claims = jwt.MapClaims{
		"userID":    userID,
		"IssuedAt":  time.Now().Unix(),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	RefreashToken, err := token.SignedString([]byte(GetEnv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return RefreashToken, nil
}
