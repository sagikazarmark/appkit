package errors

import (
	"errors"
	"testing"
)

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
