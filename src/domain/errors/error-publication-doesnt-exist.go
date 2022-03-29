package errors

import (
	"net/http"
)

type ErrorPublicationDoesntExist struct {
	ErrorCore
}

func NewErPublicationDoesntExist(err error, message string) *ErrorCore {
	return NewErrorCore(err, "/src/domain/errors/error-company-doesnt-exist.go", message, http.StatusInternalServerError)
}
