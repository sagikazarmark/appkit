package http

import (
	"net/http"

	kitxhttp "github.com/sagikazarmark/kitx/transport/http"

	"github.com/sagikazarmark/appkit/errors"
)

// DefaultProblemMatchers is a list of default ProblemMatchers.
// nolint: gochecknoglobals
var DefaultProblemMatchers = []kitxhttp.ProblemMatcher{
	kitxhttp.NewStatusProblemMatcher(http.StatusNotFound, kitxhttp.ErrorMatcherFunc(errors.IsNotFoundError)),
	kitxhttp.NewStatusProblemMatcher(http.StatusUnprocessableEntity, kitxhttp.ErrorMatcherFunc(errors.IsValidationError)),
	kitxhttp.NewStatusProblemMatcher(http.StatusBadRequest, kitxhttp.ErrorMatcherFunc(errors.IsBadRequestError)),
	kitxhttp.NewStatusProblemMatcher(http.StatusConflict, kitxhttp.ErrorMatcherFunc(errors.IsConflictError)),
}
