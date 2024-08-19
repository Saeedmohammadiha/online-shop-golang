package validation

import (
	"fmt"

	dto "github.com/OnlineShop/internal/app/dto/users"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/pkg/logger"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type IUserValidation interface {
	ValidateCreateUser(r dto.CreateUserRequest) error
	ValidateUpdateUser(r dto.UserUpdateRequest) error
}

type UserValidation struct {
	log logger.Ilogger
}

func NewUserValidator(l logger.Ilogger) IUserValidation {
	return &UserValidation{
		log: l,
	}
}

func (v *UserValidation) ValidateCreateUser(r dto.CreateUserRequest) error {

	err := validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Length(3, 20).Error("the name must be between and 20 characters")),
		validation.Field(&r.LastName, validation.Length(3, 20).Error("the name must be between and 20 characters")),
		validation.Field(&r.Email, validation.Required.Error("you must provide the email"), is.Email.Error("the email is invalid")),
		validation.Field(&r.PhoneNumber, validation.Length(11, 11).Error("phone number must be 11 character"), is.Digit.Error("phone number should be number")),
		validation.Field(&r.Password, utils.NewValidatePassword(r.Password)))
	if err != nil {
		v.log.Error("validation error accuared for login", err)
		return fmt.Errorf("%s%w", utils.ErrValidationTag, err)
	}

	v.log.Info("login validation was successful")
	return nil

}

func (v *UserValidation) ValidateUpdateUser(r dto.UserUpdateRequest) error {

	err := validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Length(3, 20).Error("the name must be between and 20 characters")),
		validation.Field(&r.LastName, validation.Length(3, 20).Error("the name must be between and 20 characters")),
		validation.Field(&r.Email, validation.Required.Error("you must provide the email"), is.Email.Error("the email is invalid")),
		validation.Field(&r.PhoneNumber, validation.Length(11, 11).Error("phone number must be 11 character"), is.Digit.Error("phone number should be number")),
		validation.Field(&r.Password, utils.NewValidatePassword(r.Password)),
	)

	if err != nil {
		v.log.Error("validation error accuared for login", err)
		return fmt.Errorf("%s%w", utils.ErrValidationTag, err)
	}

	v.log.Info("login validation was successful")
	return nil

}
