package apperrors

import "errors"

var ErrValidation = errors.New("validationError")
var ErrDatabase = errors.New("databaseError")


func HandleError()  {
	
}