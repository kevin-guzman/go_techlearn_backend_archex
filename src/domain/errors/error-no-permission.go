package errors

import (
	"net/http"
)

type ErrorUserPermission struct {
	ErrorCore
}

func NewErrorUserPermission(err error, message string) *ErrorCore {
	return NewErrorCore(err, "/src/domain/errors/error-no-permission.go", message, http.StatusUnauthorized)
}
