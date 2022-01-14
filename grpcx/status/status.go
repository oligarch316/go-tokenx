package status

import (
	"github.com/oligarch316/go-tokenx/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const unknownClassCode = codes.Internal

var knownClassCodes = map[errors.Class]codes.Code{
	errors.ClassInvalidTokenData:      codes.Unauthenticated,
	errors.ClassInvalidTokenSignature: codes.Unauthenticated,
	errors.ClassInvalidKey:            codes.Internal,
}

func FromError(err error) (*status.Status, bool) {
	if s, ok := status.FromError(err); ok {
		return s, true
	}

	if code, ok := knownClassCodes[errors.ClassFrom(err)]; ok {
		return status.New(code, err.Error()), true
	}

	return status.New(unknownClassCode, err.Error()), false
}

func Convert(err error) *status.Status {
	res, _ := FromError(err)
	return res
}
