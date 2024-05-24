package interceptor

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
}

func ValidateInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if validator, ok := req.(validator); ok {
		if err := validator.Validate(); err != nil {
			return nil, errors.New("Validate request failed: Invalid request")
		}
	}

	return handler(ctx, req)
}
