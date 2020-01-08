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
