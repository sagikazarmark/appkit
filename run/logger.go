package run

type ServeLogger interface {
	Info(msg string, fields ...map[string]interface{})
}

// LogServe returns an actor, i.e. an execute and interrupt func, that
// logs a message when a server is being started/shut down.
func LogServe(logger ServeLogger) Wrapper {
	return func(execute func() error, interrupt func(error)) (func() error, func(error)) {
		return func() error {
				logger.Info("starting server")

				return execute()
			}, func(err error) {
				logger.Info("shutting server down")

				interrupt(err)
			}
	}
}
