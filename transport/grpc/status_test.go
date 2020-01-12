package grpc

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
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
			st := converter.NewStatus(context.Background(), test.err)

			if want, have := test.expectedCode, st.Code(); want != have {
				t.Errorf("unexpected status code\nexpected: %d\nactual:   %d", want, have)
			}

			if want, have := test.err.Error(), st.Message(); want != have {
				t.Errorf("unexpected message\nexpected: %s\nactual:   %s", want, have)
			}
		})
	}
}

type validationWithViolationsStub struct{}

func (validationWithViolationsStub) Error() string {
	return "validation"
}

func (validationWithViolationsStub) Validation() bool {
	return true
}

func (validationWithViolationsStub) Violations() map[string][]string {
	return map[string][]string{
		"field": {
			"violation",
		},
	}
}

func TestDefaultStatusMatchers_ValidationWithViolations(t *testing.T) {
	converter := NewStatusConverter(WithStatusMatchers(DefaultStatusMatchers...))

	err := validationWithViolationsStub{}

	st := converter.NewStatus(context.Background(), err)

	if want, have := codes.InvalidArgument, st.Code(); want != have {
		t.Errorf("unexpected status code\nexpected: %d\nactual:   %d", want, have)
	}

	if want, have := err.Error(), st.Message(); want != have {
		t.Errorf("unexpected message\nexpected: %s\nactual:   %s", want, have)
	}

	violations, ok := st.Details()[0].(*errdetails.BadRequest)
	if !ok {
		t.Fatal("status is expected to contain violation information")
	}

	if want, have := "field", violations.GetFieldViolations()[0].GetField(); want != have {
		t.Errorf("unexpected violation field\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := "violation", violations.GetFieldViolations()[0].GetDescription(); want != have {
		t.Errorf("unexpected violation description\nexpected: %s\nactual:   %s", want, have)
	}
}
