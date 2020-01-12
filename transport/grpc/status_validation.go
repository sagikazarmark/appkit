package grpc

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appkiterrors "github.com/sagikazarmark/appkit/errors"
)

// NewValidationStatusMatcher returns a status matcher for validation errors.
// If the returned error matches the following interface, violation info gets attached to the returned status:
// 	type violationError interface {
// 		Violations() map[string][]string
// 	}
func NewValidationStatusMatcher() StatusMatcher {
	return validationStatusConverter{}
}

type validationStatusConverter struct{}

func (v validationStatusConverter) MatchError(err error) bool {
	return appkiterrors.IsValidationError(err)
}

func (v validationStatusConverter) NewStatus(_ context.Context, err error) *status.Status {
	var verr interface {
		Violations() map[string][]string
	}

	if errors.As(err, &verr) {
		st := status.New(codes.InvalidArgument, err.Error())

		br := &errdetails.BadRequest{}

		for field, violations := range verr.Violations() {
			for _, violation := range violations {
				br.FieldViolations = append(br.FieldViolations, &errdetails.BadRequest_FieldViolation{
					Field:       field,
					Description: violation,
				})
			}
		}

		st, err := st.WithDetails(br)
		if err != nil {
			// If this errored, it will always error
			// here, so better panic so we can figure
			// out why than have this silently passing.
			panic(fmt.Errorf("unexpected error attaching metadata: %w", err))
		}

		return st
	}

	return status.New(codes.InvalidArgument, err.Error())
}
