package endpoint

import (
	"context"
	"errors"
	"testing"

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
