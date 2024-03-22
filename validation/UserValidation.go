package validation

import (
	"fmt"
	"net/http"
	"unicode"

	"github.com/OnlineShop/models"
	"github.com/thedevsaddam/govalidator"
)

type Validation interface {
	Validate() error
}

type UserValidator struct {
	user *models.User
}

func NewUserValidator(u *models.User) *UserValidator {
	govalidator.AddCustomRule("password", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)

		var (
			hasMinLen  = false
			hasUpper   = false
			hasLower   = false
			hasNumber  = false
			hasSpecial = false
			err        error
		)
		if len(val) >= 7 {
			hasMinLen = true
		} else {
			err = fmt.Errorf("the %s field must have minimum 7 charachter", field)
			return err
		}
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

		if hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial {
			return nil
		} else {
			err = fmt.Errorf("the %s field must have mat least a number and an upercase and a lowercase and a symbol character", field)
			return err
		}
	})
	return &UserValidator{user: u}
}

func (v *UserValidator) Validate(r *http.Request) map[string]interface{} {

	rules := govalidator.MapData{
		"name":        []string{"required", "between:3,12"},
		"email":       []string{"required", "min:8", "max:40", "email"},
		"lastName":    []string{"min:3", "max:20"},
		"phoneNumber": []string{"digits:11"},
		"DiscountID":  []string{"numeric"},
		"password":    []string{"password"},
	}

	messages := govalidator.MapData{
		"name":        []string{"required: you must provide the name", "between:the name must be between  3 to 12 character"},
		"email":       []string{"required:you must provide email", "min:email should be minimum 8 charachter", "max:email should be maximum 40 charachter", "email: email is not valid"},
		"lastName":    []string{"min:last name should be minimum 3 charachter", "max:last name should be maximum 20 charachter"},
		"phoneNumber": []string{"digits:should be 11 numeric character"},
		"DiscountID":  []string{"should be numeric"},
		"password":    []string{"password: should be at least 7 and containe lowercase, uppercase, number and symbolic characters"},
	}

	opts := govalidator.Options{
		Request:         r,
		Rules:           rules,
		Messages:        messages,
		RequiredDefault: true, // all the field to be pass the rules
	}

	gov := govalidator.New(opts)
	e := gov.Validate()
	err := map[string]interface{}{"validationError": e}
	return err
}
