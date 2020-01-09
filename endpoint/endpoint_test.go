package endpoint

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-kit/kit/endpoint"
)

type clientErrorStub struct{}

func (clientErrorStub) Error() string {
	return "error"
}

func (clientErrorStub) ClientError() bool {
	return true
}

type errorWrapper struct {
	err error
}

func (e errorWrapper) Error() string {
	return "error"
}

func (e errorWrapper) Unwrap() error {
	return e.err
}

func TestClientErrorMiddleware(t *testing.T) {
	t.Run("client_error", func(t *testing.T) {
		origErr := errorWrapper{clientErrorStub{}}

		ep := func(ctx context.Context, request interface{}) (interface{}, error) {
			return nil, origErr
		}

		ep = ClientErrorMiddleware(ep)

		resp, err := ep(context.Background(), nil)

		if err != nil {
			t.Fatal("endpoint is NOT supposed to return an error")
		}

		failer, ok := resp.(endpoint.Failer)
		if !ok {
			t.Fatal("endpoint is supposed to return an endpoint.Failer response")
		}

		if failer.Failed() != origErr {
			t.Error("failer is supposed to return the original error")
		}
	})

	t.Run("non_client_error", func(t *testing.T) {
		origErr := errors.New("error")

		ep := func(ctx context.Context, request interface{}) (interface{}, error) {
			return nil, origErr
		}

		ep = ClientErrorMiddleware(ep)

		_, err := ep(context.Background(), nil)

		if !errors.Is(origErr, err) {
			t.Error("endpoint is NOT supposed to return an endpoint.Failer response")
		}
	})
}

// TODO: use logur test logger
// nolint: godox
type loggerStub struct {
	logs []struct {
		log    string
		fields map[string]interface{}
	}
}

func (l *loggerStub) TraceContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	var f map[string]interface{}

	if len(fields) > 0 {
		f = fields[0]
	}

	l.logs = append(l.logs, struct {
		log    string
		fields map[string]interface{}
	}{log: msg, fields: f})
}

func TestLoggingMiddleware(t *testing.T) {
	ep := func(ctx context.Context, request interface{}) (interface{}, error) {
		time.Sleep(2 * time.Millisecond)

		return nil, nil
	}

	logger := &loggerStub{}

	ep = LoggingMiddleware(logger)(ep)

	_, _ = ep(context.Background(), nil)

	if len(logger.logs) != 2 {
		t.Fatal("logger is supposed to have two messages")
	}

	if want, have := "processing request", logger.logs[0].log; want != have {
		t.Errorf("unexpected log message\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := "processing request finished", logger.logs[1].log; want != have {
		t.Errorf("unexpected log\nexpected: %s\nactual:   %s", want, have)
	}

	if logger.logs[1].fields["took"].(time.Duration) < 2*time.Millisecond {
		t.Error("the request took less than 2ms")
	}
}
