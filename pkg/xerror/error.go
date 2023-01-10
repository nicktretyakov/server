package xerror

import "fmt"

type Code string

const (
	Internal         Code = "internal"
	ProfileFailed    Code = "profile_failed"
	PermissionDenied Code = "permission_denied"
	UnAuthorized     Code = "unauthorized"
)

func New(code Code, m string) *Error {
	return &Error{
		Code:    code,
		Message: m,
	}
}

func Newf(code Code, f string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(f, args...),
	}
}

type Error struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}
