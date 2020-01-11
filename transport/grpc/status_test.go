package grpc

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
)

type notFoundStub struct{}

func (notFoundStub) Error() string {
	return "not found"
}

func (notFoundStub) NotFound() bool {
	return true
}

type validationStub struct{}

func (validationStub) Error() string {
	return "validation"
}

func (validationStub) Validation() bool {
	return true
}

type conflictStub struct{}

func (conflictStub) Error() string {
	return "conflict"
}

func (conflictStub) Conflict() bool {
	return true
}

func TestDefaultStatusMatchers(t *testing.T) {
	tests := []struct {
		err          error
		expectedCode codes.Code
	}{
		{
			err:          notFoundStub{},
			expectedCode: codes.NotFound,
		},
		{
			err:          validationStub{},
			expectedCode: codes.InvalidArgument,
		},
		{
			err:          conflictStub{},
			expectedCode: codes.FailedPrecondition,
		},
	}

	converter := NewStatusConverter(WithStatusMatchers(DefaultStatusMatchers...))

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("%d", test.expectedCode), func(t *testing.T) {
			status := converter.NewStatus(context.Background(), test.err)

			if want, have := test.expectedCode, status.Code(); want != have {
				t.Errorf("unexpected status code\nexpected: %d\nactual:   %d", want, have)
			}

			if want, have := test.err.Error(), status.Message(); want != have {
				t.Errorf("unexpected message\nexpected: %s\nactual:   %s", want, have)
			}
		})
	}
}
