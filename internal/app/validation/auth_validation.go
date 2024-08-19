package validation

import (
	"fmt"

	dto "github.com/OnlineShop/internal/app/dto/auth"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/pkg/logger"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type IAuthValidation interface {
	ValidateLogin(requestData *dto.LoginRequest) error
}

type AuthValidation struct {
	log logger.Ilogger
}

func NewAuthValidator(l logger.Ilogger) IAuthValidation {
	return &AuthValidation{
		log: l,
	}
}

func (v *AuthValidation) ValidateLogin(requestData *dto.LoginRequest) error {

	err := validation.ValidateStruct(requestData,
		validation.Field(&requestData.Email, validation.Required.Error("you must provide the email"), is.Email.Error("the email is invalid")),
		validation.Field(&requestData.Password, validation.Required.Error("you must provide the password"), utils.NewValidatePassword(requestData.Password)))

	if err != nil {
		v.log.Error("validation error accuared for login", err)
		return fmt.Errorf("%s%w", utils.ErrValidationTag, err)
	}

	v.log.Info("login validation was successful")
	return nil
}
