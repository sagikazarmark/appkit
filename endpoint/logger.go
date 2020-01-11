package endpoint

import (
	"context"
)

// Logger logs certain events of the application.
type Logger interface {
	// TraceContext logs a Trace event.
	TraceContext(ctx context.Context, msg string, fields ...map[string]interface{})
}
