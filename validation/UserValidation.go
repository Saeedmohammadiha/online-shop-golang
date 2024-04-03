package validation

import (
	"github.com/OnlineShop/dto/User"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Validation interface {
	ValidateCreateUser(r dto.CreateUserRequest) error
	ValidateUpdateUser(r dto.UserUpdateRequest) error
}

type UserValidator struct{}

func NewUserValidator() Validation {
	return &UserValidator{}
}

func (*UserValidator) ValidateCreateUser(r dto.CreateUserRequest) error {

	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Length(3, 20).Error("the name must be between and 20 characters")),
		validation.Field(&r.LastName, validation.Length(3, 20).Error("the name must be between and 20 characters")),
		validation.Field(&r.Email, validation.Required.Error("you must privide the email"), is.Email.Error("the email is invalid")),
		validation.Field(&r.PhoneNumber, validation.Length(11, 11).Error("phone number must be 11 character"), is.Digit.Error("phone number should be number")),
		validation.Field(&r.Password, NewValidatePassword(r.Password)))

}

func (*UserValidator) ValidateUpdateUser(r dto.UserUpdateRequest) error {

	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Length(3, 20).Error("the name must be between and 20 characters")),
		validation.Field(&r.LastName, validation.Length(3, 20).Error("the name must be between and 20 characters")),
		validation.Field(&r.Email, validation.Required.Error("you must privide the email"), is.Email.Error("the email is invalid")),
		validation.Field(&r.PhoneNumber, validation.Length(11, 11).Error("phone number must be 11 character"), is.Digit.Error("phone number should be number")),
		validation.Field(&r.Password, NewValidatePassword(r.Password)),
	)

}
