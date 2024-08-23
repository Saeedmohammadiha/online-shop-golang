package usecases

import (
	"context"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/auth"
	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/app/validation"
	internaljwt "github.com/OnlineShop/internal/pkg/jwt"
	"github.com/OnlineShop/internal/pkg/logger"
)

type IAuthUsecases interface {
	Login(ctx context.Context) (*dto.LoginResponse, *apperrors.AppError)
	Logout(ctx context.Context) (*dto.LogoutResponse, *apperrors.AppError)
	Refresh(ctx context.Context) (*dto.RefreshResponse, *apperrors.AppError)
}

type AuthUsecase struct {
	log        logger.Ilogger
	jwt        internaljwt.IJwtToken
	validator  validation.IAuthValidation
	repository repositories.IUserRepository
}

func NewAuthUsecase(r repositories.IUserRepository, v validation.IAuthValidation, j internaljwt.IJwtToken, l logger.Ilogger) IAuthUsecases {
	return &AuthUsecase{
		log:        l,
		validator:  v,
		repository: r,
		jwt:        j,
	}
}

func (a *AuthUsecase) Logout(ctx context.Context) (*dto.LogoutResponse, *apperrors.AppError) {
	//TODO: invalidate the tokens
	var user models.User
	err := utils.GetValueFromCtx(ctx, utils.USER, &user, a.log)
	if err != nil {
		a.log.Error("failed to get user form ctx")
		return nil, apperrors.NewAuthenticationError("failed to get the user", err)
	}
	a.log.Debug("got the user from context",
		"user", user,
	)

	user.AccessToken = ""
	user.RefreshToken = ""
	_, err = a.repository.Update(&user)
	if err != nil {
		return nil, err
	}
	a.log.Debug("deleted the tokens form db",
		"user", user,
	)
	a.log.Debug("return a success response")

	return &dto.LogoutResponse{
		Message: "you are successfully loged out",
	}, nil
}

func (a *AuthUsecase) Login(ctx context.Context) (*dto.LoginResponse, *apperrors.AppError) {

	var requestBody dto.LoginRequest
	if err := utils.GetValueFromCtx(ctx, utils.REQUEST_BODY, &requestBody, a.log); err != nil {
		return nil, err
	}
	//validate the inputs
	if err := a.validator.ValidateLogin(&requestBody); err != nil {
		return nil, err
	}

	//is there a user
	user, err := a.repository.IsUserExists(requestBody.Email)
	if err != nil {
		a.log.Error("there is no user with this email",
			"data", requestBody,
			"errorLocation", "loginUsecase",
		)
		return nil, err
	}

	//is the password correct
	ok := utils.CheckPasswordHash(&requestBody.Password, &user.Password)
	if !ok {
		a.log.Error("invalid password",
			"data", requestBody,
			"errorLocation", "loginUsecase",
		)

		return nil, apperrors.NewAuthenticationError("the password in not correct", nil)
	}

	//return token and login the user
	tokeGenerator := internaljwt.New(a.log)
	accessToken, err := tokeGenerator.GenerateAccessToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	refreshToken, err := tokeGenerator.GenerateRefreshToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	// ________________________

	responseData := dto.LoginResponse{
		User:         *user,
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}
	return &responseData, nil
}

func (a *AuthUsecase) Refresh(ctx context.Context) (*dto.RefreshResponse, *apperrors.AppError) {

	var requestBody dto.RefreshRequest
	if err := utils.GetValueFromCtx(ctx, utils.REQUEST_BODY, &requestBody, a.log); err != nil {
		// utils.SendErrorResponse(ctx, err, w, a.log)
		return nil, err
	}

	// check if the refreshToken is valid

	claims, err := a.jwt.DecodeToken(requestBody.RefreshToken)
	if err != nil {
		return nil, err
	}

	//check if the token is the same stored in the db
	user, err := a.repository.GetById(claims.UserID)
	if user.RefreshToken != requestBody.RefreshToken {
		//generate new error
		return nil, err
	}

	//generate new refresh token
	newAccessToken, err := a.jwt.GenerateAccessToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	// return the token
	return &dto.RefreshResponse{
		AccessToken: *newAccessToken,
	}, nil
}
