package run

// TemporalWorker is a Temporal worker instance.
//
// See: https://github.com/temporalio/temporal-go-sdk
type TemporalWorker interface {
	// Start starts the worker in a non-blocking fashion.
	Start() error

	// Stop cleans up any resources opened by worker.
	Stop()
}

// TemporalWorkerActor returns an actor, i.e. an execute and interrupt func, that
// terminates when the underlying worker fails or stops running.
//
// Although the Temporal Worker component has a blocking Run function,
// internally it waits for a SIGTERM signal which does not fit perfectly into run group.
//
// See https://github.com/temporalio/temporal-go-sdk/issues/94
func TemporalWorkerActor(worker TemporalWorker) (execute func() error, interrupt func(error)) {
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
