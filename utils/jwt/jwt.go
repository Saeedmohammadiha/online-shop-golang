package utils

import (
	"time"

	utils "github.com/OnlineShop/utils/env"
	"github.com/golang-jwt/jwt/v5"
)

func NewAccessToken(userID int) (string, error) {
	var claims = jwt.MapClaims{
		"userID":    userID,
		"IssuedAt":  time.Now().Unix(),
		"ExpiresAt": time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	AccessToken, err := token.SignedString([]byte(utils.GetEnv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return AccessToken, nil
}

func NewRefreshToken(userID int) (string, error) {
	var claims = jwt.MapClaims{
		"userID":    userID,
		"IssuedAt":  time.Now().Unix(),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	RefreashToken, err := token.SignedString([]byte(utils.GetEnv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return RefreashToken, nil
}
