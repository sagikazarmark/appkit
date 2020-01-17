package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/moogar0880/problems"
)

func TestNewStatusProblemMatcher(t *testing.T) {
	matcher := NewStatusProblemMatcher(http.StatusNotFound, func(err error) bool { return true })

	if !matcher.MatchError(errors.New("error")) {
		t.Error("error is supposed to be matched")
	}

	if want, have := http.StatusNotFound, matcher.Status(); want != have {
		t.Errorf("unexpected status\nexpected: %d\nactual:   %d", want, have)
	}
}

type matcherStub struct {
	err error
}

func (s matcherStub) MatchError(err error) bool {
	return s.err == err
}

type matcherConverterStub struct {
	err error
}

func (s matcherConverterStub) MatchError(err error) bool {
	return s.err == err
}

func (s matcherConverterStub) NewProblem(_ context.Context, _ error) interface{} {
	return problems.NewDetailedProblem(http.StatusServiceUnavailable, "my error")
}

type statusMatcherStub struct {
	err    error
	status int
}

func (s statusMatcherStub) MatchError(err error) bool {
	return s.err == err
}

func (s statusMatcherStub) Status() int {
	return s.status
}

type statusMatcherConverterStub struct {
	statusMatcherStub
}

func (s statusMatcherConverterStub) NewProblem(_ context.Context, _ error) interface{} {
	return problems.NewDetailedProblem(http.StatusBadRequest, "custom error")
}

func testProblemEquals(t *testing.T, problem *problems.DefaultProblem, status int, detail string) {
	t.Helper()

	if want, have := status, problem.Status; want != have {
		t.Errorf("unexpected status\nexpected: %d\nactual:   %d", want, have)
	}

	if want, have := detail, problem.Detail; want != have {
		t.Errorf("unexpected status\nexpected: %s\nactual:   %s", want, have)
	}
}

func TestProblemConverter(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		problemConverter := NewProblemConverter()

		problem := problemConverter.NewProblem(context.Background(), errors.New("error")).(*problems.DefaultProblem)

		testProblemEquals(t, problem, http.StatusInternalServerError, "something went wrong")
	})

	t.Run("matcher", func(t *testing.T) {
		err := errors.New("error")

		tests := []struct {
			options []ProblemConverterOption
			status  int
			detail  string
		}{
			{
				options: []ProblemConverterOption{
					WithProblemMatchers(statusMatcherStub{
						err:    err,
						status: http.StatusNotFound,
					}),
				},
				status: http.StatusNotFound,
				detail: "error",
			},
			{
				options: []ProblemConverterOption{
					WithProblemMatchers(statusMatcherConverterStub{
						statusMatcherStub: statusMatcherStub{
							err:    err,
							status: http.StatusNotFound,
						},
					}),
				},
				status: http.StatusBadRequest,
				detail: "custom error",
			},
			{
				options: []ProblemConverterOption{
					WithProblemMatchers(matcherStub{
						err: err,
					}),
				},
				status: http.StatusInternalServerError,
				detail: "error",
			},
			{
				options: []ProblemConverterOption{
					WithProblemMatchers(matcherConverterStub{
						err: err,
					}),
				},
				status: http.StatusServiceUnavailable,
				detail: "my error",
			},
			{ // WithProblemMatchers is supposed to append matchers to the list
				options: []ProblemConverterOption{
					WithProblemMatchers(statusMatcherStub{
						err:    err,
						status: http.StatusNotFound,
					}),
					WithProblemMatchers(statusMatcherStub{
						err:    err,
						status: http.StatusBadRequest,
					}),
				},
				status: http.StatusNotFound,
				detail: "error",
			},
		}

		for _, test := range tests {
			test := test

			t.Run("", func(t *testing.T) {
				problemConverter := NewProblemConverter(test.options...)

				problem := problemConverter.NewProblem(context.Background(), err).(*problems.DefaultProblem)

				testProblemEquals(t, problem, test.status, test.detail)
			})
		}
	})
}

func ExampleNewProblemConverter() {
	problemConverter := NewProblemConverter(
		WithProblemMatchers(
			NewStatusProblemMatcher(http.StatusNotFound, func(err error) bool { return err.Error() == "not found" }),
		),
	)

	err := errors.New("not found")

	problem := problemConverter.NewProblem(context.Background(), err).(*problems.DefaultProblem)

	fmt.Println(problem.Status, problem.Detail)

	// Output: 404 not found
}
