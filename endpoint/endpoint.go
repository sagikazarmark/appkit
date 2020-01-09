package endpoint

import (
	"context"
	"time"

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

// LoggingMiddleware logs trace information about every request.
func LoggingMiddleware(logger Logger) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.TraceContext(ctx, "processing request")

			defer func(begin time.Time) {
				logger.TraceContext(ctx, "processing request finished", map[string]interface{}{
					"took": time.Since(begin),
				})
			}(time.Now())

			return e(ctx, request)
		}
	}
}
