package errors

import (
	"errors"
)

type clientError interface {
	ClientError() bool
}

// IsClientError checks if an error should be returned to the client for processing.
// An error is considered to be a client error if it implements the following interface:
//
// 	type clientError interface {
// 		ClientError() bool
// 	}
//
// and `ClientError` returns true.
func IsClientError(err error) bool {
	var e clientError

	if errors.As(err, &e) {
		return e.ClientError()
	}

	return false
}

type notFound interface {
	NotFound() bool
}

// IsNotFoundError checks if an error is related to a resource being not found.
// An error is considered to be a NotFound error if it implements the following interface:
//
// 	type notFound interface {
// 		NotFound() bool
// 	}
//
// and `NotFound` returns true.
func IsNotFoundError(err error) bool {
	var e notFound

	if errors.As(err, &e) {
		return e.NotFound()
	}

	return false
}

type validation interface {
	Validation() bool
}

// IsValidationError checks if an error is related to a resource being not found.
// An error is considered to be a Validation error if it implements the following interface:
//
// 	type validation interface {
// 		Validation() bool
// 	}
//
// and `Validation` returns true.
func IsValidationError(err error) bool {
	var e validation

	if errors.As(err, &e) {
		return e.Validation()
	}

	return false
}
