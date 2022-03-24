package errors

import (
	"net/http"
)

type ErrorPort struct {
	ErrorCore
}

func NewErrorPort(err error) *ErrorCore {
	return NewErrorCore(err, "/src/domain/errors/error-port.go", "Error interno del servidor", http.StatusInternalServerError)
}
