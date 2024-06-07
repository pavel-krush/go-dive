package go_dive

import (
	"context"
	"strings"

	"google.golang.org/grpc"
)

type contextKey int

const (
	activityKey contextKey = iota
)

func ActivityToContext(ctx context.Context, a *Activity) context.Context {
	return context.WithValue(ctx, activityKey, a)
}

func ActivityFromContext(ctx context.Context) *Activity {
	a, _ := ctx.Value(activityKey).(*Activity)
	return a
}

func UnaryServerDiveInterceptor(master *Master) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if master == nil {
			return handler(ctx, req)
		}

		service, method := splitFullMethod(info.FullMethod)
		activity := NewUnaryGrpcActivity(service, method, req)
		ctx = ActivityToContext(ctx, activity)
		master.AddActivity(activity.id.String(), activity)
		defer master.RemoveActivity(activity.id.String())
		return handler(ctx, req)
	}
}

func splitFullMethod(fullMethod string) (service string, method string) {
	parts := strings.Split(fullMethod, "/")
	if len(parts) == 3 {
		return parts[1], parts[2]
	}
	return "", ""
}
