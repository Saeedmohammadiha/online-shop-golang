package apperrors

const ValidationErrorTag = "ValidationError"
const DatabaseErrorTag = "databaseError"
const AuthenticationErrorTag = "authenticationError"
const AuthorizationErrorTag = "authorizationError"
const InternalErrorTag = "InternalError"
const BadRequestErrorTag = "BadRequestError"

//const ErrConvertTag = "convertJsonError"
//const ErrInvalidPasswordTag = "invalidPassword"
//const ErrTokenGenerationTag = "TokenGeneration"
//const ErrTokenDecodingTag = "TokenDecode"
//const ErrTokenInvalidTag = "InvalidToken"

type AppError struct {
	Message string
	Tag     string
	Err     error
}

func (a *AppError) Error() string {
	return a.Message
}

func NewAuthenticationError(m string, err error) *AppError {
	return &AppError{
		Message: m,
		Err:     err,
		Tag:     AuthenticationErrorTag,
	}
}

func NewValidationError(m string, err error) *AppError {
	return &AppError{
		Message: m,
		Err:     err,
		Tag:     ValidationErrorTag,
	}
}

func NewInternalError(m string, err error) *AppError {
	return &AppError{
		Message: m,
		Err:     err,
		Tag:     InternalErrorTag,
	}
}
func NewDatabaseError(m string, err error) *AppError {
	return &AppError{
		Message: m,
		Err:     err,
		Tag:     DatabaseErrorTag,
	}
}
func NewBadRequestError(m string, err error) *AppError {
	return &AppError{
		Message: m,
		Err:     err,
		Tag:     BadRequestErrorTag,
	}
}
func NewAuthorizationError(m string, err error) *AppError {
	return &AppError{
		Message: m,
		Err:     err,
		Tag:     AuthorizationErrorTag,
	}
}
