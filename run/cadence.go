package run

// CadenceWorker is an Uber Cadence worker instance.
type CadenceWorker interface {
	// Start starts the worker in a non-blocking fashion.
	Start() error

	// Stop cleans up any resources opened by worker.
	Stop()
}

// CadenceWorkerRun returns an actor, i.e. an execute and interrupt func, that
// terminates when the underlying worker fails.
//
// Although the Cadence Worker component has a blocking Run function,
// internally it waits for a SIGTERM signal which does not fit perfectly into run group.
//
// See https://github.com/uber-go/cadence-client/issues/642
func CadenceWorkerRun(worker CadenceWorker) (execute func() error, interrupt func(error)) {
	var closeCh = make(chan struct{})

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
