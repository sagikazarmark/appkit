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

func TestClientErrorMatcher(t *testing.T) {
	t.Run("ClientError", func(t *testing.T) {
		if !ClientErrorMatcher().MatchError(clientErrorStub{}) {
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
				if ClientErrorMatcher().MatchError(err) {
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

func TestNotFoundErrorMatcher(t *testing.T) {
	t.Run("NotFoundError", func(t *testing.T) {
		if !NotFoundErrorMatcher().MatchError(notFoundStub{}) {
			t.Error("error is supposed to be a NotFoundError")
		}
	})

	t.Run("NonNotFoundError", func(t *testing.T) {
		tests := []error{
			errors.New("error"),
			nonNotFoundStub{},
		}

		for _, err := range tests {
			err := err

			t.Run("", func(t *testing.T) {
				if NotFoundErrorMatcher().MatchError(err) {
					t.Error("error is NOT supposed to be a NotFoundError")
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

func TestValidationErrorMatcher(t *testing.T) {
	t.Run("ValidationError", func(t *testing.T) {
		if !ValidationErrorMatcher().MatchError(validationStub{}) {
			t.Error("error is supposed to be a ValidationError")
		}
	})

	t.Run("NonValidationError", func(t *testing.T) {
		tests := []error{
			errors.New("error"),
			nonValidationStub{},
		}

		for _, err := range tests {
			err := err

			t.Run("", func(t *testing.T) {
				if ValidationErrorMatcher().MatchError(err) {
					t.Error("error is NOT supposed to be a ValidationError")
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

func TestBadRequestErrorMatcher(t *testing.T) {
	t.Run("BadRequestError", func(t *testing.T) {
		if !BadRequestErrorMatcher().MatchError(badRequestStub{}) {
			t.Error("error is supposed to be a BadRequestError")
		}
	})

	t.Run("NonBadRequestError", func(t *testing.T) {
		tests := []error{
			errors.New("error"),
			nonBadRequestStub{},
		}

		for _, err := range tests {
			err := err

			t.Run("", func(t *testing.T) {
				if BadRequestErrorMatcher().MatchError(err) {
					t.Error("error is NOT supposed to be a BadRequestError")
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

func TestConflictErrorMatcher(t *testing.T) {
	t.Run("ConflictError", func(t *testing.T) {
		if !ConflictErrorMatcher().MatchError(conflictStub{}) {
			t.Error("error is supposed to be a ConflictError")
		}
	})

	t.Run("NonConflictError", func(t *testing.T) {
		tests := []error{
			errors.New("error"),
			nonConflictStub{},
		}

		for _, err := range tests {
			err := err

			t.Run("", func(t *testing.T) {
				if ConflictErrorMatcher().MatchError(err) {
					t.Error("error is NOT supposed to be a ConflictError")
				}
			})
		}
	})
}
