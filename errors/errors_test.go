package errors

import (
	"errors"
	"testing"
)

type serviceErrorStub struct{}

func (serviceErrorStub) Error() string {
	return ""
}

func (serviceErrorStub) ServiceError() bool {
	return true
}

type nonServiceErrorStub struct{}

func (c nonServiceErrorStub) Error() string {
	return ""
}

func (c nonServiceErrorStub) ServiceError() bool {
	return false
}

func TestIsServiceError(t *testing.T) {
	t.Run("ServiceError", func(t *testing.T) {
		if !IsServiceError(serviceErrorStub{}) {
			t.Error("error is supposed to be a ServiceError")
		}
	})

	t.Run("NonServiceError", func(t *testing.T) {
		tests := []error{
			errors.New("error"),
			nonServiceErrorStub{},
		}

		for _, err := range tests {
			err := err

			t.Run("", func(t *testing.T) {
				if IsServiceError(err) {
					t.Error("error is NOT supposed to be a ServiceError")
				}
			})
		}
	})
}

type clientErrorStub struct{}

func (clientErrorStub) Error() string {
	return ""
}

func (clientErrorStub) ClientError() bool {
	return true
}

type nonClientErrorStub struct{}

func (c nonClientErrorStub) Error() string {
	return ""
}

func (c nonClientErrorStub) ClientError() bool {
	return false
}

func TestIsClientError(t *testing.T) {
	t.Run("ClientError", func(t *testing.T) {
		if !IsClientError(clientErrorStub{}) {
			t.Error("error is supposed to be a ClientError")
		}
	})

	t.Run("NonClientError", func(t *testing.T) {
		tests := []error{
			errors.New("error"),
			nonClientErrorStub{},
		}

		for _, err := range tests {
			err := err

			t.Run("", func(t *testing.T) {
				if IsClientError(err) {
					t.Error("error is NOT supposed to be a ClientError")
				}
			})
		}
	})
}

type notFoundStub struct{}

func (notFoundStub) Error() string {
	return ""
}

func (notFoundStub) NotFound() bool {
	return true
}

type nonNotFoundStub struct{}

func (c nonNotFoundStub) Error() string {
	return ""
}

func (c nonNotFoundStub) NotFound() bool {
	return false
}

func TestIsNotFoundError(t *testing.T) {
	t.Run("NotFound", func(t *testing.T) {
		if !IsNotFoundError(notFoundStub{}) {
			t.Error("error is supposed to be a NotFound error")
		}
	})

	t.Run("NonNotFound", func(t *testing.T) {
		tests := []error{
			errors.New("error"),
			nonNotFoundStub{},
		}

		for _, err := range tests {
			err := err

			t.Run("", func(t *testing.T) {
				if IsNotFoundError(err) {
					t.Error("error is NOT supposed to be a NotFound error")
				}
			})
		}
	})
}

type validationStub struct{}

func (validationStub) Error() string {
	return ""
}

func (validationStub) Validation() bool {
	return true
}

type nonValidationStub struct{}

func (c nonValidationStub) Error() string {
	return ""
}

func (c nonValidationStub) Validation() bool {
	return false
}

func TestIsValidationError(t *testing.T) {
	t.Run("Validation", func(t *testing.T) {
		if !IsValidationError(validationStub{}) {
			t.Error("error is supposed to be a Validation error")
		}
	})

	t.Run("NonValidation", func(t *testing.T) {
		tests := []error{
			errors.New("error"),
			nonValidationStub{},
		}

		for _, err := range tests {
			err := err

			t.Run("", func(t *testing.T) {
				if IsValidationError(err) {
					t.Error("error is NOT supposed to be a Validation error")
				}
			})
		}
	})
}

type badRequestStub struct{}

func (badRequestStub) Error() string {
	return ""
}

func (badRequestStub) BadRequest() bool {
	return true
}

type nonBadRequestStub struct{}

func (c nonBadRequestStub) Error() string {
	return ""
}

func (c nonBadRequestStub) BadRequest() bool {
	return false
}

func TestIsBadRequestError(t *testing.T) {
	t.Run("BadRequest", func(t *testing.T) {
		if !IsBadRequestError(badRequestStub{}) {
			t.Error("error is supposed to be a BadRequest error")
		}
	})

	t.Run("NonBadRequest", func(t *testing.T) {
		tests := []error{
			errors.New("error"),
			nonBadRequestStub{},
		}

		for _, err := range tests {
			err := err

			t.Run("", func(t *testing.T) {
				if IsBadRequestError(err) {
					t.Error("error is NOT supposed to be a BadRequest error")
				}
			})
		}
	})
}

type conflictStub struct{}

func (conflictStub) Error() string {
	return ""
}

func (conflictStub) Conflict() bool {
	return true
}

type nonConflictStub struct{}

func (c nonConflictStub) Error() string {
	return ""
}

func (c nonConflictStub) Conflict() bool {
	return false
}

func TestIsConflictError(t *testing.T) {
	t.Run("Conflict", func(t *testing.T) {
		if !IsConflictError(conflictStub{}) {
			t.Error("error is supposed to be a Conflict error")
		}
	})

	t.Run("NonConflict", func(t *testing.T) {
		tests := []error{
			errors.New("error"),
			nonConflictStub{},
		}

		for _, err := range tests {
			err := err

			t.Run("", func(t *testing.T) {
				if IsConflictError(err) {
					t.Error("error is NOT supposed to be a Conflict error")
				}
			})
		}
	})
}
