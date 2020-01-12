package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/moogar0880/problems"
)

// ProblemConverter converts an error to a RFC-7807 Problem.
//
// See details at https://tools.ietf.org/html/rfc7807
type ProblemConverter interface {
	// NewProblem creates a new RFC-7807 Problem from an error.
	// A problem can be any structure that marshals to an RFC-7807 compatible JSON/XML structure.
	NewProblem(ctx context.Context, err error) interface{}
}

// StatusProblem is the interface describing a problem with an associated Status code.
type StatusProblem interface {
	ProblemStatus() int
}

// ProblemMatcher matches an error.
// A ProblemMatcher usually also implements one of the following interfaces:
//
// - StatusProblemMatcher to indicate an HTTP status code for an error
// - ProblemConverter if a matched error requires special conversion logic
type ProblemMatcher interface {
	// MatchError evaluates the predefined set of conditions for err.
	MatchError(err error) bool
}

// StatusProblemMatcher matches an error and returns the appropriate status code for it.
type StatusProblemMatcher interface {
	ProblemMatcher

	// Status returns the HTTP status code.
	Status() int
}

type statusProblemMatcher struct {
	errorMatcher ErrorMatcher
	status       int
}

// ErrorMatcher checks if an error matches a certain condition.
type ErrorMatcher func(err error) bool

// NewStatusProblemMatcher returns a new StatusProblemMatcher.
func NewStatusProblemMatcher(status int, errorMatcher ErrorMatcher) StatusProblemMatcher {
	return statusProblemMatcher{
		errorMatcher: errorMatcher,
		status:       status,
	}
}

func (m statusProblemMatcher) MatchError(err error) bool {
	return m.errorMatcher(err)
}

func (m statusProblemMatcher) Status() int {
	return m.status
}

// StatusProblemConverter converts an error to a RFC-7807 Problem.
//
// See details at https://tools.ietf.org/html/rfc7807
type StatusProblemConverter interface {
	// NewStatusProblem creates a new RFC-7807 Problem with a status code.
	NewStatusProblem(ctx context.Context, status int, err error) StatusProblem
}

type defaultProblemConverter struct{}

func (c defaultProblemConverter) NewProblem(_ context.Context, err error) interface{} {
	return problems.NewDetailedProblem(http.StatusInternalServerError, err.Error())
}

func (c defaultProblemConverter) NewStatusProblem(_ context.Context, status int, err error) StatusProblem {
	return problems.NewDetailedProblem(status, err.Error())
}

type problemConverter struct {
	matchers []ProblemMatcher

	problemConverter       ProblemConverter
	statusProblemConverter StatusProblemConverter
}

// ProblemConverterOption configures a ProblemConverter using the functional options paradigm
// popularized by Rob Pike and Dave Cheney.
// If you're unfamiliar with this style, see:
// - https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
// - https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis.
type ProblemConverterOption interface {
	apply(c *problemConverter)
}

type problemConverterOptionFunc func(*problemConverter)

func (f problemConverterOptionFunc) apply(c *problemConverter) { f(c) }

// WithProblemMatchers configures a ProblemConverter to match errors.
// By default an empty problem is created.
// If no matchers match an error (or no matchers are configured) an HTTP 500 problem is returned.
//
// If a matcher also implements ProblemConverter it is used instead of the builtin ProblemConverter
// for creating the problem.
//
// If a matcher also implements StatusProblemMatcher
// the builtin StatusProblemConverter is used for creating the problem.
func WithProblemMatchers(matchers ...ProblemMatcher) ProblemConverterOption {
	return problemConverterOptionFunc(func(c *problemConverter) {
		c.matchers = matchers
	})
}

// WithProblemConverter configures a ProblemConverter.
func WithProblemConverter(converter ProblemConverter) ProblemConverterOption {
	return problemConverterOptionFunc(func(c *problemConverter) {
		c.problemConverter = converter
	})
}

// WithStatusProblemConverter configures a StatusProblemConverter.
func WithStatusProblemConverter(converter StatusProblemConverter) ProblemConverterOption {
	return problemConverterOptionFunc(func(c *problemConverter) {
		c.statusProblemConverter = converter
	})
}

// NewProblemConverter returns a new ProblemConverter implementation.
func NewProblemConverter(opts ...ProblemConverterOption) ProblemConverter {
	c := problemConverter{}

	for _, opt := range opts {
		opt.apply(&c)
	}

	if c.problemConverter == nil {
		c.problemConverter = defaultProblemConverter{}
	}

	if c.statusProblemConverter == nil {
		if spc, ok := c.problemConverter.(StatusProblemConverter); ok {
			c.statusProblemConverter = spc
		} else {
			c.statusProblemConverter = defaultProblemConverter{}
		}
	}

	return c
}

func (c problemConverter) NewProblem(ctx context.Context, err error) interface{} {
	for _, matcher := range c.matchers {
		if matcher.MatchError(err) {
			if converter, ok := matcher.(ProblemConverter); ok {
				return converter.NewProblem(ctx, err)
			}

			if statusMatcher, ok := matcher.(StatusProblemMatcher); ok {
				return c.statusProblemConverter.NewStatusProblem(ctx, statusMatcher.Status(), err)
			}

			return c.problemConverter.NewProblem(ctx, err)
		}
	}

	return c.statusProblemConverter.NewStatusProblem(
		ctx,
		http.StatusInternalServerError,
		errors.New("something went wrong"),
	)
}

// NewProblemConverter returns a new ProblemConverter implementation populated with default problem matchers.
func NewDefaultProblemConverter(opts ...ProblemConverterOption) ProblemConverter {
	opts = append([]ProblemConverterOption{WithProblemMatchers(DefaultProblemMatchers...)}, opts...)

	return NewProblemConverter(opts...)
}
