package grpc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewStatusCodeMatcher(t *testing.T) {
	matcher := NewStatusCodeMatcher(codes.NotFound, func(err error) bool { return true })

	if !matcher.MatchError(errors.New("error")) {
		t.Error("error is supposed to be matched")
	}

	if want, have := codes.NotFound, matcher.Code(); want != have {
		t.Errorf("unexpected code\nexpected: %d\nactual:   %d", want, have)
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

func (s matcherConverterStub) NewStatus(_ context.Context, _ error) *status.Status {
	return status.New(codes.Unavailable, "my error")
}

type statusMatcherStub struct {
	err  error
	code codes.Code
}

func (s statusMatcherStub) MatchError(err error) bool {
	return s.err == err
}

func (s statusMatcherStub) Code() codes.Code {
	return s.code
}

type statusMatcherConverterStub struct {
	statusMatcherStub
}

func (s statusMatcherConverterStub) NewStatus(_ context.Context, _ error) *status.Status {
	return status.New(codes.InvalidArgument, "custom error")
}

func testStatusEquals(t *testing.T, s *status.Status, code codes.Code, message string) {
	t.Helper()

	if want, have := code, s.Code(); want != have {
		t.Errorf("unexpected code\nexpected: %d\nactual:   %d", want, have)
	}

	if want, have := message, s.Message(); want != have {
		t.Errorf("unexpected code\nexpected: %s\nactual:   %s", want, have)
	}
}

func TestStatusConverter(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		statusConverter := NewStatusConverter()

		s := statusConverter.NewStatus(context.Background(), errors.New("error"))

		testStatusEquals(t, s, codes.Internal, "something went wrong")
	})

	t.Run("matcher", func(t *testing.T) {
		err := errors.New("error")

		tests := []struct {
			options []StatusConverterOption
			code    codes.Code
			message string
		}{
			{
				options: []StatusConverterOption{
					WithStatusMatchers(statusMatcherStub{
						err:  err,
						code: codes.NotFound,
					}),
				},
				code:    codes.NotFound,
				message: "error",
			},
			{
				options: []StatusConverterOption{
					WithStatusMatchers(statusMatcherConverterStub{
						statusMatcherStub: statusMatcherStub{
							err:  err,
							code: http.StatusNotFound,
						},
					}),
				},
				code:    codes.InvalidArgument,
				message: "custom error",
			},
			{
				options: []StatusConverterOption{
					WithStatusMatchers(matcherStub{
						err: err,
					}),
				},
				code:    codes.Internal,
				message: "error",
			},
			{
				options: []StatusConverterOption{
					WithStatusMatchers(matcherConverterStub{
						err: err,
					}),
				},
				code:    codes.Unavailable,
				message: "my error",
			},
			{ // WithStatusMatchers is supposed to append matchers to the list
				options: []StatusConverterOption{
					WithStatusMatchers(statusMatcherStub{
						err:  err,
						code: codes.NotFound,
					}),
					WithStatusMatchers(statusMatcherStub{
						err:  err,
						code: codes.InvalidArgument,
					}),
				},
				code:    codes.NotFound,
				message: "error",
			},
		}

		for _, test := range tests {
			test := test

			t.Run("", func(t *testing.T) {
				statusConverter := NewStatusConverter(test.options...)

				s := statusConverter.NewStatus(context.Background(), err)

				testStatusEquals(t, s, test.code, test.message)
			})
		}
	})
}

func ExampleNewStatusConverter() {
	statusConverter := NewStatusConverter(
		WithStatusMatchers(
			NewStatusCodeMatcher(codes.NotFound, func(err error) bool { return err.Error() == "not found" }),
		),
	)

	err := errors.New("not found")

	s := statusConverter.NewStatus(context.Background(), err)

	fmt.Println(s.Code(), s.Message())

	// Output: NotFound not found
}
