package http

import (
	"net/http"

	"github.com/sagikazarmark/appkit/errors"
)

// DefaultProblemMatchers is a list of default ProblemMatchers.
// nolint: gochecknoglobals
var DefaultProblemMatchers = []ProblemMatcher{
	NewStatusProblemMatcher(http.StatusNotFound, errors.IsNotFoundError),
	NewValidationWithViolationsProblemMatcher(),
	NewStatusProblemMatcher(http.StatusUnprocessableEntity, errors.IsValidationError),
	NewStatusProblemMatcher(http.StatusBadRequest, errors.IsBadRequestError),
	NewStatusProblemMatcher(http.StatusConflict, errors.IsConflictError),
}
