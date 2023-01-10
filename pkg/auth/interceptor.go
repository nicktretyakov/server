package auth

import (
	"context"
	"strings"

	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"be/internal/model"
)

var UserCtxKey = &contextKey{name: "user"} //nolint:gochecknoglobals

type contextKey struct {
	name string
}

func Interceptor(authService IAuth) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		if strings.HasPrefix(strings.TrimPrefix(info.FullMethod, "/"), "booking.Auth") {
			return handler(ctx, req)
		}

		return grpcAuth.UnaryServerInterceptor(authFunc(authService))(ctx, req, info, handler)
	}
}

func FromContext(ctx context.Context) *model.User {
	user, ok := ctx.Value(UserCtxKey).(*model.User)
	if !ok || user == nil {
		return nil
	}

	return user
}

func authFunc(authService IAuth) grpcAuth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		token, err := fromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}

		user, _, xerr := authService.ParseToken(ctx, token)
		if xerr != nil {
			return nil, status.New(codes.Unauthenticated, xerr.Error()).Err()
		}

		return context.WithValue(ctx, UserCtxKey, user), nil
	}
}

func fromMD(ctx context.Context, expectedScheme string) (string, error) {
	val := metautils.ExtractIncoming(ctx).Get("authorization")
	if val == "" {
		return "", status.New(codes.Unauthenticated, "expected bearer token").Err()
	}

	splits := strings.SplitN(val, " ", 2)
	if len(splits) < 2 {
		return "", status.New(codes.Unauthenticated, "invalid authorization header").Err()
	}

	if !strings.EqualFold(splits[0], expectedScheme) {
		return "", status.New(codes.Unauthenticated, "invalid authorization header").Err()
	}

	return splits[1], nil
}
