package http

import (
	"net/http"

	kitxhttp "github.com/sagikazarmark/kitx/transport/http"

	"github.com/sagikazarmark/appkit/errors"
)

// DefaultProblemMatchers is a list of default ProblemMatchers.
// nolint: gochecknoglobals
var DefaultProblemMatchers = []kitxhttp.ProblemMatcher{
	kitxhttp.NewStatusProblemMatcher(http.StatusNotFound, errors.NotFoundErrorMatcher()),
	kitxhttp.NewStatusProblemMatcher(http.StatusUnprocessableEntity, errors.ValidationErrorMatcher()),
	kitxhttp.NewStatusProblemMatcher(http.StatusBadRequest, errors.BadRequestErrorMatcher()),
	kitxhttp.NewStatusProblemMatcher(http.StatusConflict, errors.ConflictErrorMatcher()),
}
