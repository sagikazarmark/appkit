package errors

import (
	"errors"
)

// ErrorMatcher checks if an error matches a certain condition.
type ErrorMatcher interface {
	// MatchError checks if err matches a certain condition.
	MatchError(err error) bool
}

// errorMatcherFunc turns a plain function into an ErrorMatcher if it's definition matches the interface.
type errorMatcherFunc func(err error) bool

// MatchError calls the underlying function to check if err matches a certain condition.
func (fn errorMatcherFunc) MatchError(err error) bool {
	return fn(err)
}

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

// ClientErrorMatcher returns an error matcher that matches client errors.
// An error is considered to be a client error if it implements the following interface:
//
// 	type clientError interface {
// 		ClientError() bool
// 	}
//
// and `ClientError` returns true.
func ClientErrorMatcher() ErrorMatcher {
	return errorMatcherFunc(IsClientError)
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

// NotFoundErrorMatcher returns an error matcher that matches not found errors.
// An error is considered to be a NotFound error if it implements the following interface:
//
// 	type notFound interface {
// 		NotFound() bool
// 	}
//
// and `NotFound` returns true.
func NotFoundErrorMatcher() ErrorMatcher {
	return errorMatcherFunc(IsNotFoundError)
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

// ValidationErrorMatcher returns an error matcher that matches not found errors.
// An error is considered to be a Validation error if it implements the following interface:
//
// 	type validation interface {
// 		Validation() bool
// 	}
//
// and `Validation` returns true.
func ValidationErrorMatcher() ErrorMatcher {
	return errorMatcherFunc(IsValidationError)
}

type badRequest interface {
	BadRequest() bool
}

// IsBadRequestError checks if an error is related to a resource being not found.
// An error is considered to be a BadRequest error if it implements the following interface:
//
// 	type badRequest interface {
// 		BadRequest() bool
// 	}
//
// and `BadRequest` returns true.
func IsBadRequestError(err error) bool {
	var e badRequest

	if errors.As(err, &e) {
		return e.BadRequest()
	}

	return false
}

// BadRequestErrorMatcher returns an error matcher that matches not found errors.
// An error is considered to be a BadRequest error if it implements the following interface:
//
// 	type badRequest interface {
// 		BadRequest() bool
// 	}
//
// and `BadRequest` returns true.
func BadRequestErrorMatcher() ErrorMatcher {
	return errorMatcherFunc(IsBadRequestError)
}

type conflict interface {
	Conflict() bool
}

// IsConflictError checks if an error is related to a resource being not found.
// An error is considered to be a Conflict error if it implements the following interface:
//
// 	type conflict interface {
// 		Conflict() bool
// 	}
//
// and `Conflict` returns true.
func IsConflictError(err error) bool {
	var e conflict

	if errors.As(err, &e) {
		return e.Conflict()
	}

	return false
}

// ConflictErrorMatcher returns an error matcher that matches not found errors.
// An error is considered to be a Conflict error if it implements the following interface:
//
// 	type conflict interface {
// 		Conflict() bool
// 	}
//
// and `Conflict` returns true.
func ConflictErrorMatcher() ErrorMatcher {
	return errorMatcherFunc(IsConflictError)
}
