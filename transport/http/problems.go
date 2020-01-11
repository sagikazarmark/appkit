package http

import (
	"net/http"

	"github.com/sagikazarmark/appkit/errors"
)

// DefaultProblemMatchers is a list of default ProblemMatchers.
// nolint: gochecknoglobals
var DefaultProblemMatchers = []ProblemMatcher{
	NewStatusProblemMatcher(http.StatusNotFound, errors.NotFoundErrorMatcher()),
	NewStatusProblemMatcher(http.StatusUnprocessableEntity, errors.ValidationErrorMatcher()),
	NewStatusProblemMatcher(http.StatusBadRequest, errors.BadRequestErrorMatcher()),
	NewStatusProblemMatcher(http.StatusConflict, errors.ConflictErrorMatcher()),
}
