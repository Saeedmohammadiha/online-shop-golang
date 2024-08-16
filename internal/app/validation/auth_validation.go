package validation

import (
	dto "github.com/OnlineShop/internal/app/dto/auth"
	"github.com/OnlineShop/internal/app/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type AuthValidation interface {
	ValidateLogin(r dto.LoginRequest) error
}

type AuthhV struct{}

func NewAuthValidator() AuthValidation {
	return &AuthhV{}
}

func (*AuthhV) ValidateLogin(r dto.LoginRequest) error {

	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required.Error("you must provide the email"), is.Email.Error("the email is invalid")),
		validation.Field(&r.Password, validation.Required.Error("you must provide the password"), utils.NewValidatePassword(r.Password)))
}
