package run

// CadenceWorker is an Uber Cadence worker instance.
//
// See https://godoc.org/go.uber.org/cadence/internal#Worker
type CadenceWorker interface {
	// Start starts the worker in a non-blocking fashion.
	Start() error

	// Stop cleans up any resources opened by worker.
	Stop()
}

// CadenceWorkerActor returns an actor, i.e. an execute and interrupt func, that
// terminates when the underlying worker fails or stops running.
//
// Although the Cadence Worker component has a blocking Run function,
// internally it waits for a SIGTERM signal which does not fit perfectly into run group.
//
// See https://github.com/uber-go/cadence-client/issues/642
func CadenceWorkerActor(worker CadenceWorker) (execute func() error, interrupt func(error)) {
	closeCh := make(chan struct{})

	return func() error {
			err := worker.Start()
			if err != nil {
				return err
			}

			<-closeCh

			return nil
		}, func(error) {
			worker.Stop()
			close(closeCh)
		}
}

// CadenceWorkerActor returns an actor, i.e. an execute and interrupt func, that
// terminates when the underlying worker fails or stops running.

// Deprecated: use CadenceWorkerActor instead.
func CadenceWorkerRun(worker CadenceWorker) (execute func() error, interrupt func(error)) {
	return CadenceWorkerActor(worker)
}
