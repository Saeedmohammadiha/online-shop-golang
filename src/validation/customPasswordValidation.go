package validation

import (
	"errors"
	"fmt"
	"unicode"
)

type ValidatePassword struct {
	value string
}

func NewValidatePassword(v string) *ValidatePassword {
	return &ValidatePassword{value: v}
}

func (vp *ValidatePassword) Validate(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		return errors.New("input is not a string")
	}

	if len(val) < 7 {
		return fmt.Errorf("the %s field must have a minimum of 7 characters", vp.value)
	}

	var (
		hasUpper, hasLower, hasNumber, hasSpecial bool
	)

	for _, char := range val {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return fmt.Errorf("the %s field must have at least one uppercase, one lowercase, one number, and one special character", vp.value)
	}

	return nil
}
