package errors

import (
	"errors"
	"testing"
)

type clientErrorStub struct{}

func (clientErrorStub) Error() string {
	return "error"
}

func (clientErrorStub) ClientError() bool {
	return true
}

type nonClientErrorStub struct{}

func (c nonClientErrorStub) Error() string {
	return "error"
}

func (c nonClientErrorStub) ClientError() bool {
	return false
}

func TestIsClientError(t *testing.T) {
	t.Run("client_error", func(t *testing.T) {
		if !IsClientError(clientErrorStub{}) {
			t.Error("error is supposed to be a ClientError")
		}
	})

	t.Run("non_client_error", func(t *testing.T) {
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
