package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/moogar0880/problems"

	appkiterrors "github.com/sagikazarmark/appkit/errors"
)

// NewValidationWithViolationsProblemMatcher returns a problem matcher for validation errors that contain violations.
// If the returned error matches the following interface, a special validation problem is returned by NewProblem:
// 	type violationError interface {
// 		Violations() map[string][]string
// 	}
func NewValidationWithViolationsProblemMatcher() ProblemMatcher {
	return validationWithViolationsProblemMatcher{}
}

type violationError interface {
	Violations() map[string][]string
}

type validationWithViolationsProblemMatcher struct{}

func (v validationWithViolationsProblemMatcher) MatchError(err error) bool {
	var verr violationError

	return appkiterrors.IsValidationError(err) && errors.As(err, &verr)
}

func (v validationWithViolationsProblemMatcher) NewProblem(_ context.Context, err error) interface{} {
	var verr violationError

	if errors.As(err, &verr) {
		return NewValidationProblem(err.Error(), verr.Violations())
	}

	return problems.NewDetailedProblem(http.StatusUnprocessableEntity, err.Error())
}

// ValidationProblem describes an RFC-7807 problem with validation violations.
type ValidationProblem struct {
	*problems.DefaultProblem

	Violations map[string][]string `json:"violations"`
}

// NewValidationProblem returns a problem with details and validation errors.
func NewValidationProblem(details string, violations map[string][]string) *ValidationProblem {
	return &ValidationProblem{
		DefaultProblem: problems.NewDetailedProblem(http.StatusUnprocessableEntity, details),
		Violations:     violations,
	}
}
