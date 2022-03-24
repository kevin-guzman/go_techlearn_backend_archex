package errors

import (
	"net/http"
)

type ErrorUserDoentExist struct {
	ErrorCore
}

func NewErrorUserDoentExist(err error, message string) *ErrorCore {
	return NewErrorCore(err, "/src/domain/errors/error-user-doesnt-exist.go", message, http.StatusInternalServerError)
}
