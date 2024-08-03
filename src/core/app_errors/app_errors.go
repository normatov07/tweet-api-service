package app_errors

import "fmt"

type ApplicationError struct {
	code      int
	errorText string
}

func (e ApplicationError) Error() string {
	return e.errorText
}

func (e ApplicationError) Code() int {
	return e.code
}
func NewAppErr(code int, errorDetails ...string) *ApplicationError {
	err := &ApplicationError{code: code, errorText: getCodeText(code)}
	for _, detail := range errorDetails {
		err.errorText = fmt.Sprintf("%s. %s", err.errorText, detail)
	}
	return err
}

func NewFromErr(code int, errors ...error) *ApplicationError {
	err := &ApplicationError{code: code, errorText: getCodeText(code)}
	for _, errorItem := range errors {
		(*err).errorText += fmt.Sprintf("%s. %s", err.errorText, errorItem.Error())
	}
	return err
}

func getCodeText(code int) string {
	switch code {
	case USER_NOT_FOUND:
		return "user not found"
	case SERVER_ERROR:
		return "server error"
	case LOGIN_ERROR:
		return "login or password is incorrect"
	case LOGIN_UNIQUE:
		return "the login has been already taken"
	default:
		return ""
	}
}

const (
	USER_NOT_FOUND = -100
	LOGIN_ERROR    = -101
	LOGIN_UNIQUE   = -102
	SERVER_ERROR   = -500
)
