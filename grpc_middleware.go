package dive

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

func UnaryServerInterceptor(master *Master) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if master == nil {
			return handler(ctx, req)
		}

		service, method := splitFullMethod(info.FullMethod)
		activity := NewUnaryGrpcActivity(service, method, req)
		ctx = ActivityToContext(ctx, activity)
		master.AddActivity(activity.id, activity)
		defer master.RemoveActivity(activity.id)
		return handler(ctx, req)
	}
}

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedStream) Context() context.Context {
	return w.ctx
}

func StreamServerInterceptor(master *Master) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if master == nil {
			return handler(srv, stream)
		}

		service, method := splitFullMethod(info.FullMethod)

		activity := NewStreamGrpcActivity(service, method, nil)
		ctx := ActivityToContext(stream.Context(), activity)
		master.AddActivity(activity.id, activity)
		defer master.RemoveActivity(activity.id)

		ss := &wrappedStream{
			ServerStream: stream,
			ctx:          ctx,
		}

		return handler(srv, ss)
	}
}

func splitFullMethod(fullMethod string) (service string, method string) {
	parts := strings.Split(fullMethod, "/")
	if len(parts) == 3 {
		return parts[1], parts[2]
	}
	return "", ""
}
