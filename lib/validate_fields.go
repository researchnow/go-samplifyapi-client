package samplify

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

const (
	// InvalidField indicates that the input field is invalid
	InvalidField = "INVALID_FIELD"
)

// ValidateError ...
type ValidateError interface {
	FieldID() string
	Error() string
	Code() string
}

// FieldError ...
type FieldError struct {
	Err validator.FieldError
}

// FieldID ...
func (f *FieldError) FieldID() string {
	return f.Err.Field()
}

// Error ...
func (f *FieldError) Error() string {
	result := fmt.Sprintf("invalid value for " + f.Err.Field() + " allowed " + f.Err.Tag() + " " + f.Err.Param())
	return result
}

// Code ...
func (f *FieldError) Code() string {
	return InvalidField
}
