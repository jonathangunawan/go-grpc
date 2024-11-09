package interceptor

import (
	"context"
	"fmt"
	"strings"

	"github.com/jonathangunawan/go-grpc/constant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Interceptor struct{}

// if we need only to accept interceptor for certain path
// then you need to create conditional by yourself to handle this
// you can get the path from *grpc.UnaryServerInfo and field FullMethod
func (i Interceptor) UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if _, ok := unaryMethods[info.FullMethod]; ok {
		// get header from metadata context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf(constant.ErrValidationEmptyAuth)
		}

		if !valid(md["authorization"]) {
			return nil, fmt.Errorf(constant.ErrValidationInvalidToken)
		}
	}

	// this is the real handler you will be executed
	// so the position of your interceptor code before or after the handler will matters
	m, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}

	return m, err
}

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")

	// Perform token validation here
	// For this example we only check if the token is exact match with hardcoded value

	return token == "some-token"
}
