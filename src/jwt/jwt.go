package JwtFactory

import (
	"os"

	"github.com/OnlineShop/src/models"
	"github.com/golang-jwt/jwt/v5"
)

type Token interface {
	GenerateAccessToken(userID int, roleIDs []models.Role) (string, error)
	GenerateRefreshToken(userID int, roleIDs []models.Role) (string, error)
}

type JWTClaim struct {
	UserID int `json:"userID"`
	RoleID int `json:"roleID"`
	jwt.RegisteredClaims
}

type Jwt struct{}

func New() Token {
	return &Jwt{}
}

func (*Jwt) GenerateAccessToken(userID int, roleIDs []models.Role) (string, error) {
	
	var claims = &JWTClaim{
		UserID: userID,
		RoleID: int(roleIDs[0].ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	AccessToken, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return AccessToken, nil
}

func (*Jwt) GenerateRefreshToken(userID int,  roleIDs []models.Role) (string, error) {
	var claims = &JWTClaim{
		UserID: userID,
		RoleID:  int(roleIDs[0].ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	RefreashToken, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return RefreashToken, nil
}
