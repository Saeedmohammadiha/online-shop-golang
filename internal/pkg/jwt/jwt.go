package internaljwt

import (
	"os"

	"github.com/OnlineShop/internal/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
)

type IJwtToken interface {
	GenerateAccessToken(userID int) (*string, error)
	GenerateRefreshToken(userID int) (*string, error)
}

type JWTClaim struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}

type JwtToken struct {
	log logger.Ilogger
}

func New(l logger.Ilogger) IJwtToken {
	return &JwtToken{
		log: l,
	}
}

func (j *JwtToken) GenerateAccessToken(userID int) (*string, error) {

	var claims = &JWTClaim{
		UserID: userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	AccessToken, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		j.log.Debug("failed to generate access token",
			"error", err,
			"user", userID,
			"claim", claims,
		)
		return nil, err

	}
	j.log.Debug("a new access token is generated",
		"token", AccessToken,
		"user", userID,
		"claim", claims,
	)
	return &AccessToken, nil
}

func (j *JwtToken) GenerateRefreshToken(userID int) (*string, error) {
	var claims = &JWTClaim{
		UserID: userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	RefreshToken, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		j.log.Debug("failed to generate refresh token",
			"error", err,
			"user", userID,
			"claim", claims,
		)
		return nil, err
	}

	j.log.Debug("a new refresh token is generated",
		"token", RefreshToken,
		"user", userID,
		"claim", claims,
	)
	return &RefreshToken, nil
}
