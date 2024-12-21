package calc

import "errors"

var (
	ErrValidationError = errors.New("validation error")
	ErrDivisionByZero  = errors.New("division by zero")
)
