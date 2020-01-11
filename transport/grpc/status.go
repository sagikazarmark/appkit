package grpc

import (
	"google.golang.org/grpc/codes"

	"github.com/sagikazarmark/appkit/errors"
)

// DefaultStatusMatchers is a list of default StatusMatchers.
// nolint: gochecknoglobals
var DefaultStatusMatchers = []StatusMatcher{
	NewStatusCodeMatcher(codes.NotFound, errors.NotFoundErrorMatcher()),
	NewStatusCodeMatcher(codes.InvalidArgument, errors.ValidationErrorMatcher()),
	NewStatusCodeMatcher(codes.FailedPrecondition, errors.ConflictErrorMatcher()),
}
