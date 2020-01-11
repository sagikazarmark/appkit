package grpc

import (
	kitxgrpc "github.com/sagikazarmark/kitx/transport/grpc"
	"google.golang.org/grpc/codes"

	"github.com/sagikazarmark/appkit/errors"
)

// DefaultStatusMatchers is a list of default StatusMatchers.
// nolint: gochecknoglobals
var DefaultStatusMatchers = []StatusMatcher{
	NewStatusCodeMatcher(codes.NotFound, kitxgrpc.ErrorMatcherFunc(errors.IsNotFoundError)),
	NewStatusCodeMatcher(codes.InvalidArgument, kitxgrpc.ErrorMatcherFunc(errors.IsValidationError)),
	NewStatusCodeMatcher(codes.FailedPrecondition, kitxgrpc.ErrorMatcherFunc(errors.IsConflictError)),
}
