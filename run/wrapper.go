package run

// Wrapper is an actor that wraps another actor.
// It is useful for creating middleware.
type Wrapper func(execute func() error, interrupt func(error)) (func() error, func(error))
