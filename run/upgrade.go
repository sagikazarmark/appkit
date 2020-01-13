package run

import (
	"context"
)

// Upgrader handles zero downtime upgrades and passing files between processes.
//
// Upgrader is based on https://github.com/cloudflare/tableflip.
type Upgrader interface {
	Ready() error
	Exit() <-chan struct{}
	Stop()
}

// GracefulRestart returns an actor, i.e. an execute and interrupt func, that
// terminates when graceful restart is initiated and the child process
// signals to be ready, or the parent context is canceled.
func GracefulRestart(ctx context.Context, upg Upgrader) (execute func() error, interrupt func(error)) {
	ctx, cancel := context.WithCancel(ctx)

	return func() error {
			// Tell the parent we are ready
			err := upg.Ready()
			if err != nil {
				return err
			}

			select {
			case <-upg.Exit(): // Wait for child to be ready (or application shutdown)
				return nil

			case <-ctx.Done():
				return ctx.Err()
			}
		}, func(error) {
			cancel()
			upg.Stop()
		}
}
