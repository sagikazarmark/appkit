package run

import (
	"net"
)

// GRPCServer is a gRPC server.
type GRPCServer interface {
	Serve(lis net.Listener) error
	GracefulStop()
}

// GRPCServe returns an actor, i.e. an execute and interrupt func, that
// terminates when the underlying gRPC server fails.
func GRPCServe(server GRPCServer, lis net.Listener) (execute func() error, interrupt func(error)) {
	return func() error {
			return server.Serve(lis)
		}, func(error) {
			server.GracefulStop()
		}
}
