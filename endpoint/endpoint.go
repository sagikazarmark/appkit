package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"

	"github.com/sagikazarmark/appkit/errors"
)

// errorMatcherFunc turns a plain function into an ErrorMatcher if it's definition matches the interface.
type errorMatcherFunc func(err error) bool

// MatchError calls the underlying function to check if err matches a certain condition.
func (fn errorMatcherFunc) MatchError(err error) bool {
	return fn(err)
}

// ClientErrorMiddleware checks returned errors of the next endpoint.
// Errors matching the client error criteria get wrapped in an endpoint.Failer response.
// An error is considered to be a client error if it implements the following interface:
//
// 	type clientError interface {
// 		ClientError() bool
// 	}
//
// and `ClientError` returns true.
func ClientErrorMiddleware(e endpoint.Endpoint) endpoint.Endpoint {
	return kitxendpoint.FailerMiddleware(errorMatcherFunc(errors.IsClientError))(e)
}
