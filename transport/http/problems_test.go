package http

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/moogar0880/problems"
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

type badRequestStub struct{}

func (badRequestStub) Error() string {
	return "bad request"
}

func (badRequestStub) BadRequest() bool {
	return true
}

type conflictStub struct{}

func (conflictStub) Error() string {
	return "conflict"
}

func (conflictStub) Conflict() bool {
	return true
}

func TestDefaultProblemMatchers(t *testing.T) {
	tests := []struct {
		err            error
		expectedStatus int
	}{
		{
			err:            notFoundStub{},
			expectedStatus: http.StatusNotFound,
		},
		{
			err:            validationStub{},
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			err:            badRequestStub{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			err:            conflictStub{},
			expectedStatus: http.StatusConflict,
		},
	}

	converter := NewProblemConverter(WithProblemMatchers(DefaultProblemMatchers...))

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("%d", test.expectedStatus), func(t *testing.T) {
			problem := converter.NewProblem(context.Background(), test.err).(*problems.DefaultProblem)

			if want, have := test.expectedStatus, problem.Status; want != have {
				t.Errorf("unexpected status\nexpected: %d\nactual:   %d", want, have)
			}

			if want, have := test.err.Error(), problem.Detail; want != have {
				t.Errorf("unexpected detail\nexpected: %s\nactual:   %s", want, have)
			}
		})
	}
}
