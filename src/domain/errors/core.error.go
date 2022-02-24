package errors

import (
	"errors"
	"fmt"
)

type ErrorCore struct {
	Trace, Message string
	Err            error
}

var NewErrorCore = func(err error, trace, message string) *ErrorCore {
	fmt.Println("\nError has courred in \n", trace, "\nThe error is", err)
	return &ErrorCore{
		Err:     err,
		Trace:   trace,
		Message: message,
	}
}

func (e *ErrorCore) PublicError() error {
	return errors.New(e.Message)
}
