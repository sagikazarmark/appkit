package run

import (
	"context"
	"net"
	"time"
)

// HTTPServer is an HTTP server.
type HTTPServer interface {
	Serve(lis net.Listener) error
	Shutdown(ctx context.Context) error
}

// HTTPServe returns an actor, i.e. an execute and interrupt func, that
// terminates when the underlying HTTP server fails.
//
// It accepts a timeout parameter for the graceful shutdown. 0 means no timeout.
func HTTPServe(
	server HTTPServer,
	lis net.Listener,
	shutdownTimeout time.Duration,
) (execute func() error, interrupt func(error)) {
	return func() error {
			return server.Serve(lis)
		}, func(error) {
			ctx := context.Background()
			if shutdownTimeout > 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, shutdownTimeout)
				defer cancel()
			}

			_ = server.Shutdown(ctx)
		}
}
