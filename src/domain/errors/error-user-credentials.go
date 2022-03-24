package errors

import (
	"net/http"
)

type ErrorUserCredentials struct {
	ErrorCore
}

func NewErrorUserCredentials(err error) *ErrorCore {
	return NewErrorCore(err, "/src/domain/errors/error-user-credentials.go", "Creadenciales inválidas", http.StatusUnauthorized)
}
