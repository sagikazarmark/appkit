package endpoint

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/sagikazarmark/appkit/errors"
)

type failer struct {
	err error
}

func (f failer) Failed() error {
	return f.err
}

// ServiceErrorMiddleware checks returned errors of the subsequent endpoint.
// Errors matching the client error criteria get wrapped in an endpoint.Failer response.
// An error is considered to be a client error if it implements the following interface:
//
//	type serviceError interface {
//		ServiceError() bool
//	}
//
// and `ServiceError` returns true.
func ServiceErrorMiddleware(e endpoint.Endpoint) endpoint.Endpoint {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			resp, err := e(ctx, request)
			if err != nil && errors.IsServiceError(err) {
				return failer{err}, nil
			}

			return resp, err
		}
	}(e)
}

// ClientErrorMiddleware checks returned errors of the subsequent endpoint.
// Errors matching the client error criteria get wrapped in an endpoint.Failer response.
// An error is considered to be a client error if it implements the following interface:
//
//	type clientError interface {
//		ClientError() bool
//	}
//
// and `ClientError` returns true.
//
// Deprecated: use ServiceErrorMiddleware instead.
func ClientErrorMiddleware(e endpoint.Endpoint) endpoint.Endpoint {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			resp, err := e(ctx, request)
			if err != nil && errors.IsClientError(err) {
				return failer{err}, nil
			}

			return resp, err
		}
	}(e)
}

// LoggingMiddleware logs trace information about every request
// (beginning of the request, processing time).
//
// The logger might extract additional information from the context
// (correlation ID, operation name, etc).
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
