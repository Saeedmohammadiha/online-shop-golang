package usecases

import (
	"fmt"

	dto "github.com/OnlineShop/internal/app/dto/auth"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/app/validation"
	internaljwt "github.com/OnlineShop/internal/pkg/jwt"
	"github.com/OnlineShop/internal/pkg/logger"
)

type IAuthUsecases interface {
	Login(data *dto.LoginRequest) (*dto.LoginResponse, error)
}

type AuthUsecase struct {
	log        logger.Ilogger
	validator  validation.IAuthValidation
	repository repositories.IUserRepository
}

func NewAuthUsecase(r repositories.IUserRepository, v validation.IAuthValidation, l logger.Ilogger) IAuthUsecases {
	return &AuthUsecase{
		log:        l,
		validator:  v,
		repository: r,
	}
}

func (a *AuthUsecase) Login(data *dto.LoginRequest) (*dto.LoginResponse, error) {

	//validate the inputs
	if err := a.validator.ValidateLogin(data); err != nil {
		return nil, err
	}

	//is there a user
	user, err := a.repository.IsUserExists(data.Email)
	if err != nil {
		a.log.Error("there is no user with this email",
			"data", data,
			"errorLocation", "loginUsecase",
		)
		return nil, err
	}

	//is the password correct
	ok := utils.CheckPasswordHash(&data.Password, &user.Password)
	if !ok {
		a.log.Error("invalid password",
			"data", data,
			"errorLocation", "loginUsecase",
		)

		return nil, fmt.Errorf("%s%w", utils.ErrInvalidPasswordTag, err)
	}

	//return token and login the user
	tokeGenerator := internaljwt.New(a.log)
	accessToken, err := tokeGenerator.GenerateAccessToken(int(user.ID))
	if err != nil {
		return nil, fmt.Errorf("%s%w", utils.ErrTokenGenerationTag, err)
	}

	refreshToken, err := tokeGenerator.GenerateRefreshToken(int(user.ID))
	if err != nil {
		return nil, fmt.Errorf("%s%w", utils.ErrTokenGenerationTag, err)
	}
	// ________________________

	responseData := dto.LoginResponse{
		User:         *user,
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}
	return &responseData, nil
}
