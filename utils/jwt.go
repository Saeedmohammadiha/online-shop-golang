package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

type Token interface {
	NewAccessToken(userID int, roleID int) (string, error)
	NewRefreshToken(userID int, roleID int) (string, error)
}

type JWTClaim struct {
	UserID int `json:"userID"`
	RoleID int `json:"roleID"`
	jwt.RegisteredClaims
}

type Jwt struct{}

func NewAuth() Token {
	return &Jwt{}
}

func (*Jwt) NewAccessToken(userID int, roleID int) (string, error) {
	var claims = &JWTClaim{
		UserID: userID,
		RoleID: roleID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	AccessToken, err := token.SignedString([]byte(GetEnv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return AccessToken, nil
}

func (*Jwt) NewRefreshToken(userID int, roleID int) (string, error) {
	var claims = &JWTClaim{
		UserID: userID,
		RoleID: roleID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	RefreashToken, err := token.SignedString([]byte(GetEnv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return RefreashToken, nil
}
