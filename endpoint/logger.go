package endpoint

import (
	"context"
)

type Logger interface {
	TraceContext(ctx context.Context, msg string, fields ...map[string]interface{})
}
