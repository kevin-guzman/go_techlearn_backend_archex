package errors

import (
	"net/http"
)

type ErrorCompanyDoesntExist struct {
	ErrorCore
}

func NewErrorCompanyDoesntExist(err error, message string) *ErrorCore {
	return NewErrorCore(err, "/src/domain/errors/error-company-doesnt-exist.go", message, http.StatusInternalServerError)
}
