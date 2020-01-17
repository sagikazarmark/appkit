package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StatusConverter converts an error to gRPC Status.
type StatusConverter interface {
	// NewStatus creates a new gRPC Status from an error.
	NewStatus(ctx context.Context, err error) *status.Status
}

// StatusMatcher matches an error.
// A StatusMatcher usually also implements one of the following interfaces:
//
// - StatusCodeMatcher to indicate a gRPC status code for an error
// - StatusConverter if a matched error requires special conversion logic
type StatusMatcher interface {
	// MatchError evaluates the predefined set of conditions for err.
	MatchError(err error) bool
}

// StatusCodeMatcher matches an error and returns the appropriate status code for it.
type StatusCodeMatcher interface {
	StatusMatcher

	// Code returns the gRPC status code.
	Code() codes.Code
}

type statusMatcher struct {
	errorMatcher ErrorMatcher
	code         codes.Code
}

// ErrorMatcher checks if an error matches a certain condition.
type ErrorMatcher func(err error) bool

// NewStatusCodeMatcher returns a new StatusCodeMatcher.
func NewStatusCodeMatcher(code codes.Code, errorMatcher ErrorMatcher) StatusCodeMatcher {
	return statusMatcher{
		errorMatcher: errorMatcher,
		code:         code,
	}
}

func (m statusMatcher) MatchError(err error) bool {
	return m.errorMatcher(err)
}

func (m statusMatcher) Code() codes.Code {
	return m.code
}

// StatusCodeConverter converts an error to a gRPC status.
type StatusCodeConverter interface {
	// NewStatusWithCode creates a new gRPC status with a code from an error.
	NewStatusWithCode(ctx context.Context, code codes.Code, err error) *status.Status
}

type defaultStatusConverter struct{}

func (c defaultStatusConverter) NewStatus(_ context.Context, err error) *status.Status {
	return status.New(codes.Internal, err.Error())
}

func (c defaultStatusConverter) NewStatusWithCode(_ context.Context, code codes.Code, err error) *status.Status {
	return status.New(code, err.Error())
}

type statusConverter struct {
	matchers []StatusMatcher

	statusConverter     StatusConverter
	statusCodeConverter StatusCodeConverter
}

// StatusConverterOption configures a StatusConverter using the functional options paradigm
// popularized by Rob Pike and Dave Cheney.
// If you're unfamiliar with this style, see:
// - https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
// - https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis.
type StatusConverterOption interface {
	apply(c *statusConverter)
}

type statusConverterOptionFunc func(*statusConverter)

func (f statusConverterOptionFunc) apply(c *statusConverter) { f(c) }

// WithStatusMatchers configures a StatusConverter to match errors.
// Matchers are appended to the existing list of matchers.
// If no matchers match the error (or no matchers are configured) a status with Internal code is returned.
//
// If a matcher also implements StatusConverter it is used instead of the builtin StatusConverter
// for creating the status.
//
// If a matchers also implements StatusCodeMatcher
// the builtin StatusCodeConverter is used for creating the status.
func WithStatusMatchers(matchers ...StatusMatcher) StatusConverterOption {
	return statusConverterOptionFunc(func(c *statusConverter) {
		c.matchers = append(c.matchers, matchers...)
	})
}

// SetStatusMatchers configures a StatusConverter to match errors.
// Matchers override the existing list of matchers.
// If no matchers match the error (or no matchers are configured) a status with Internal code is returned.
//
// If a matcher also implements StatusConverter it is used instead of the builtin StatusConverter
// for creating the status.
//
// If a matchers also implements StatusCodeMatcher
// the builtin StatusCodeConverter is used for creating the status.
func SetStatusMatchers(matchers ...StatusMatcher) StatusConverterOption {
	return statusConverterOptionFunc(func(c *statusConverter) {
		c.matchers = matchers
	})
}

// WithStatusConverter configures a StatusConverter.
func WithStatusConverter(converter StatusConverter) StatusConverterOption {
	return statusConverterOptionFunc(func(c *statusConverter) {
		c.statusConverter = converter
	})
}

// WithStatusCodeConverter configures a StatusCodeConverter.
func WithStatusCodeConverter(converter StatusCodeConverter) StatusConverterOption {
	return statusConverterOptionFunc(func(c *statusConverter) {
		c.statusCodeConverter = converter
	})
}

// NewStatusConverter returns a new StatusConverter implementation.
func NewStatusConverter(opts ...StatusConverterOption) StatusConverter {
	c := statusConverter{}

	for _, opt := range opts {
		opt.apply(&c)
	}

	if c.statusConverter == nil {
		c.statusConverter = defaultStatusConverter{}
	}

	if c.statusCodeConverter == nil {
		if spc, ok := c.statusConverter.(StatusCodeConverter); ok {
			c.statusCodeConverter = spc
		} else {
			c.statusCodeConverter = defaultStatusConverter{}
		}
	}

	return c
}

func (c statusConverter) NewStatus(ctx context.Context, err error) *status.Status {
	for _, matcher := range c.matchers {
		if matcher.MatchError(err) {
			if converter, ok := matcher.(StatusConverter); ok {
				return converter.NewStatus(ctx, err)
			}

			if statusMatcher, ok := matcher.(StatusCodeMatcher); ok {
				return c.statusCodeConverter.NewStatusWithCode(ctx, statusMatcher.Code(), err)
			}

			return c.statusConverter.NewStatus(ctx, err)
		}
	}

	return c.statusCodeConverter.NewStatusWithCode(
		ctx,
		codes.Internal,
		errors.New("something went wrong"),
	)
}

// NewStatusConverter returns a new StatusConverter implementation populated with default status matchers.
func NewDefaultStatusConverter(opts ...StatusConverterOption) StatusConverter {
	opts = append([]StatusConverterOption{WithStatusMatchers(DefaultStatusMatchers...)}, opts...)

	return NewStatusConverter(opts...)
}
