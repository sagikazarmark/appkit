package errors

import (
	"errors"
)

type serviceError interface {
	ServiceError() bool
}

// IsServiceError checks if an error should be returned to the client for processing.
// An error is considered to be a client error if it implements the following interface:
//
// 	type serviceError interface {
// 		ServiceError() bool
// 	}
//
// and `ServiceError` returns true.
func IsServiceError(err error) bool {
	var e serviceError

	if errors.As(err, &e) {
		return e.ServiceError()
	}

	return false
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
//
// Deprecated: use ServiceError instead.
func IsClientError(err error) bool {
	var e clientError

	return errors.As(err, &e) && e.ClientError()
}

type notFound interface {
	NotFound() bool
}

// IsNotFoundError checks if an error is related to a resource not being found.
// An error is considered to be a NotFound error if it implements the following interface:
//
// 	type notFound interface {
// 		NotFound() bool
// 	}
//
// and `NotFound` returns true.
func IsNotFoundError(err error) bool {
	var e notFound

	return errors.As(err, &e) && e.NotFound()
}

type validation interface {
	Validation() bool
}

// IsValidationError checks if an error is related to a resource or request being invalid.
// An error is considered to be a Validation error if it implements the following interface:
//
// 	type validation interface {
// 		Validation() bool
// 	}
//
// and `Validation` returns true.
func IsValidationError(err error) bool {
	var e validation

	return errors.As(err, &e) && e.Validation()
}

type badRequest interface {
	BadRequest() bool
}

// IsBadRequestError checks if an error is related to a bad request being made.
// An error is considered to be a BadRequest error if it implements the following interface:
//
// 	type badRequest interface {
// 		BadRequest() bool
// 	}
//
// and `BadRequest` returns true.
func IsBadRequestError(err error) bool {
	var e badRequest

	return errors.As(err, &e) && e.BadRequest()
}

type conflict interface {
	Conflict() bool
}

// IsConflictError checks if an error is related to a resource conflict.
// An error is considered to be a Conflict error if it implements the following interface:
//
// 	type conflict interface {
// 		Conflict() bool
// 	}
//
// and `Conflict` returns true.
func IsConflictError(err error) bool {
	var e conflict

	return errors.As(err, &e) && e.Conflict()
}
