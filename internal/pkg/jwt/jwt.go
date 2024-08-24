package internaljwt

import (
	"os"
	"time"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	"github.com/OnlineShop/internal/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
)

type IJwtToken interface {
	GenerateAccessToken(userID int) (*string, *apperrors.AppError)
	GenerateRefreshToken(userID int) (*string, *apperrors.AppError)
	DecodeToken(tokenString string) (*JWTClaim, *apperrors.AppError)
}

type JWTClaim struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}

type JwtToken struct {
	log logger.Ilogger
}

func New() IJwtToken {
	l := logger.Logger()
	l.Debug(" new jwt module is created")
	return &JwtToken{
		log: l,
	}
}

func (j *JwtToken) GenerateAccessToken(userID int) (*string, *apperrors.AppError) {

	var claims = &JWTClaim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), // Access token expires in 15 minutes
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	AccessToken, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		j.log.Debug("failed to generate access token",
			"error", err,
			"user", userID,
			"claim", claims,
		)
		return nil, apperrors.NewInternalError("failed to generate access token", err)

	}
	j.log.Debug("a new access token is generated",
		"token", AccessToken,
		"user", userID,
		"claim", claims,
	)
	return &AccessToken, nil
}

func (j *JwtToken) GenerateRefreshToken(userID int) (*string, *apperrors.AppError) {
	var claims = &JWTClaim{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)), // Access token expires in 15 minutes
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	RefreshToken, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		j.log.Debug("failed to generate refresh token",
			"error", err,
			"user", userID,
			"claim", claims,
		)
		return nil, apperrors.NewInternalError("failed to generate refresh token ", err)
	}

	j.log.Debug("a new refresh token is generated",
		"token", RefreshToken,
		"user", userID,
		"claim", claims,
	)
	return &RefreshToken, nil
}

func (j *JwtToken) DecodeToken(tokenString string) (*JWTClaim, *apperrors.AppError) {
	claims := &JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		j.log.Error("failed to parse the token ",
			"token", tokenString,
		)
		return nil, apperrors.NewAuthenticationError("failed to decode token", err)
	}

	if token.Valid {
		j.log.Debug("the token is valid and parsed successfully",
			"token", tokenString,
			"decoded token", token,
			"claims", claims,
		)
		return claims, nil

	} else {
		j.log.Error("the token is not valid",
			"token", tokenString,
			"decoded token", token,
		)
		return nil, apperrors.NewAuthenticationError("the token in not valid", err)
	}

}
