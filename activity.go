package go_dive

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	diveContracts "github.com/pavel-krush/dive/api/go/v1"
)

type Type int

const (
	ActivityTypeUnknown Type = iota
	ActivityTypeBackground
	ActivityTypeGrpcUnary
	ActivityTypeGrpcStream
)

type Activity struct {
	mu          sync.RWMutex
	id          uuid.UUID
	t           Type
	startedAt   time.Time
	req         any
	props       map[string]any
	grpcService string
	grpcMethod  string
}

func newActivity(t Type, req any) *Activity {
	return &Activity{
		id:        uuid.New(),
		t:         t,
		startedAt: time.Now(),
		req:       req,
		props:     make(map[string]any),
	}
}

func NewBackgroundActivity(req any) *Activity {
	return newActivity(ActivityTypeBackground, req)
}

func NewUnaryGrpcActivity(service string, method string, req any) *Activity {
	ret := newActivity(ActivityTypeGrpcUnary, req)
	ret.grpcService = service
	ret.grpcMethod = method
	return ret
}

func NewStreamGrpcActivity(service string, method string, req any) *Activity {
	ret := newActivity(ActivityTypeGrpcStream, req)
	ret.grpcService = service
	ret.grpcMethod = method
	return ret
}

func (a *Activity) SetProps(props map[string]any) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for k, v := range props {
		a.props[k] = v
	}
}

func (a *Activity) SetProp(key string, value any) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.props[key] = value
}

func (a *Activity) ToProto() *diveContracts.Activity {
	var err error
	var reqPb *anypb.Any
	if pm, ok := a.req.(proto.Message); ok {
		reqPb, err = anypb.New(pm)
		if err != nil {
			tmp, _ := structpb.NewStruct(map[string]any{"error": "failed to convert request to anyPb"})
			reqPb, _ = anypb.New(tmp)
		}
	} else {
		tmp, _ := structpb.NewStruct(map[string]any{"error": "request is not proto.Message"})
		reqPb, _ = anypb.New(tmp)
	}

	propsPb, err := structpb.NewStruct(a.props)
	if err != nil {
		propsPb, _ = structpb.NewStruct(map[string]any{"error": "failed to convert props to structPb"})
	}

	ret := &diveContracts.Activity{
		Type:        a.t.ToProto(),
		Id:          a.id.String(),
		StartedAt:   timestamppb.New(a.startedAt),
		Duration:    durationpb.New(time.Since(a.startedAt)),
		Request:     reqPb,
		Props:       propsPb,
		GrpcService: a.grpcService,
		GrpcMethod:  a.grpcMethod,
	}

	return ret
}

func (t Type) ToProto() diveContracts.Activity_Type {
	switch t {
	case ActivityTypeBackground:
		return diveContracts.Activity_TYPE_BACKGROUND
	case ActivityTypeGrpcUnary:
		return diveContracts.Activity_TYPE_GRPC_UNARY
	case ActivityTypeGrpcStream:
		return diveContracts.Activity_TYPE_GRPC_STREAM
	case ActivityTypeUnknown:
		fallthrough
	default:
		return diveContracts.Activity_TYPE_UNKNOWN
	}
}
