package grpc

import (
	"google.golang.org/grpc/codes"

	"github.com/sagikazarmark/appkit/errors"
)

// DefaultStatusMatchers is a list of default StatusMatchers.
// nolint: gochecknoglobals
var DefaultStatusMatchers = []StatusMatcher{
	NewStatusCodeMatcher(codes.NotFound, errors.IsNotFoundError),
	NewStatusCodeMatcher(codes.InvalidArgument, errors.IsValidationError),
	NewStatusCodeMatcher(codes.FailedPrecondition, errors.IsConflictError),
}
