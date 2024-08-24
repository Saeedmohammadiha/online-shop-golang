package validation

import (
	"context"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/auth"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/pkg/logger"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type IAuthValidation interface {
	ValidateLogin(ctx context.Context,requestData *dto.LoginRequest) *apperrors.AppError
}

type AuthValidation struct {
	log logger.Ilogger
}

func NewAuthValidation() IAuthValidation {
	l := logger.Logger()
	l.Debug("new auth validation is created")
	return &AuthValidation{
		log: l,
	}
}

func (v *AuthValidation) ValidateLogin(ctx context.Context,requestData *dto.LoginRequest) *apperrors.AppError {

	err := validation.ValidateStruct(requestData,
		validation.Field(&requestData.Email, validation.Required.Error("you must provide the email"), is.Email.Error("the email is invalid")),
		validation.Field(&requestData.Password, validation.Required.Error("you must provide the password"), utils.NewValidatePassword(requestData.Password)))

	if err != nil {
		v.log.Error("validation error accuared for login", err)
		return apperrors.NewValidationError("the inputs are not valid", err)
	}

	v.log.Info("login validation was successful")
	return nil
}
