package errcode

import (
	"fmt"
)

// Error is a generic type of error
type Error struct {
	Code  string // code is the type of the error.
	error        // err is the error message, human friendly.
}

// Error codes
const (
	NotFound     = "not-found"
	InvalidArg   = "invalid-arg"
	Internal     = "internal"
	Unauthorized = "unauthorized"
)

// Add create a new error with an error code added to it.
func Add(code string, err error) *Error {
	return &Error{
		Code:  code,
		error: err,
	}
}

// Of return the code of a error if it is coded, otherwise return ""
func Of(err error) string {
	if codedErr, ok := err.(*Error); ok {
		return codedErr.Code
	}
	return ""
}

// Errorf create an Error with a code.
func Errorf(code string, f string, args ...interface{}) *Error {
	return Add(code, fmt.Errorf(f, args...))
}

// AltErrorf replace the message of an Error, keep code the same.
func AltErrorf(e *Error, msg string, args ...interface{}) *Error {
	return Errorf(e.Code, msg, args...)
}

// NotFoundf returns a not-found error.
func NotFoundf(f string, args ...interface{}) *Error {
	return Errorf(NotFound, f, args...)
}

// InvalidArgf returns an error caused by invalid argument.
func InvalidArgf(f string, args ...interface{}) *Error {
	return Errorf(InvalidArg, f, args...)
}

// Internalf returns an internal error.
func Internalf(f string, args ...interface{}) *Error {
	return Errorf(Internal, f, args...)
}

// Unauthorizedf returns an error caused by unauthrozied request.
func Unauthorizedf(f string, args ...interface{}) *Error {
	return Errorf(Unauthorized, f, args...)
}