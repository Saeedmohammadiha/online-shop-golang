package validation

import (
	//	"fmt"
	"net/http"
	"regexp"
	//	"unicode"

	"github.com/OnlineShop/models"
	"github.com/go-ozzo/ozzo-validation"
)

type Validation interface {
	ValidateCreateUser() error
	ValidateUpdateUser() error
}

type UserValidator struct {
	user *models.User
}

func NewUserValidator(u *models.User) *UserValidator {
	return &UserValidator{user: u}
}

func (v *UserValidator) ValidateCreateUser(r *http.Request) error {

	return validation.ValidateStruct(&v.user,
		validation.Field(&v.user.Name, validation.Length(3, 20)),
		validation.Field(&v.user.LastName, validation.Length(3, 20)),
		validation.Field(&v.user.Email, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"))),
		validation.Field(&v.user.PasswordHash, validation.Match(regexp.MustCompile("^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)[a-zA-Z\\d]{8,}$"))),
		validation.Field(&v.user.DiscountID, validation.Match(regexp.MustCompile("[0-9]"))),
		validation.Field(&v.user.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{11}$"))),
	)

	// rules := govalidator.MapData{
	// 	"name":        []string{"required", "between:3,12"},
	// 	"email":       []string{"required", "min:8", "max:40", "email"},
	// 	"lastName":    []string{"min:3", "max:20"},
	// 	"phoneNumber": []string{"digits:11"},
	// 	"DiscountID":  []string{"numeric"},
	// 	"password":    []string{"password"},
	// }

}

// func(field string, rule string, message string, value interface{}) error {
// 	val := value.(string)

// 	var (
// 		hasMinLen  = false
// 		hasUpper   = false
// 		hasLower   = false
// 		hasNumber  = false
// 		hasSpecial = false
// 		err        error
// 	)
// 	if len(val) >= 7 {
// 		hasMinLen = true
// 	} else {
// 		err = fmt.Errorf("the %s field must have minimum 7 charachter", field)
// 		return err
// 	}
// 	for _, char := range val {
// 		switch {
// 		case unicode.IsUpper(char):
// 			hasUpper = true
// 		case unicode.IsLower(char):
// 			hasLower = true
// 		case unicode.IsNumber(char):
// 			hasNumber = true
// 		case unicode.IsPunct(char) || unicode.IsSymbol(char):
// 			hasSpecial = true
// 		}
// 	}

// 	if hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial {
// 		return nil
// 	} else {
// 		err = fmt.Errorf("the %s field must have mat least a number and an upercase and a lowercase and a symbol character", field)
// 		return err
// 	}
// }

// messages := govalidator.MapData{
// 	"name":        []string{"required: you must provide the name", "between:the name must be between  3 to 12 character"},
// 	"email":       []string{"required:you must provide email", "min:email should be minimum 8 charachter", "max:email should be maximum 40 charachter", "email: email is not valid"},
// 	"lastName":    []string{"min:last name should be minimum 3 charachter", "max:last name should be maximum 20 charachter"},
// 	"phoneNumber": []string{"digits:should be 11 numeric character"},
// 	"DiscountID":  []string{"should be numeric"},
// 	"password":    []string{"password: should be at least 7 and containe lowercase, uppercase, number and symbolic characters"},
// }
