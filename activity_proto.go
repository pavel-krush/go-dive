package dive

import (
	"time"

	diveContracts "github.com/pavel-krush/go-dive/api/go/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func anyToAnyPb(val any) *anypb.Any {
	var err error
	var ret *anypb.Any
	if pm, ok := val.(proto.Message); ok {
		ret, err = anypb.New(pm)
		if err != nil {
			tmp, _ := structpb.NewStruct(map[string]any{"error": "failed to convert to anyPb"})
			ret, _ = anypb.New(tmp)
		}
	} else {
		tmp, _ := structpb.NewStruct(map[string]any{"error": "not proto.Message"})
		ret, _ = anypb.New(tmp)
	}
	return ret
}

func ActivityToProto(a *Activity) *diveContracts.Activity {
	ret := &diveContracts.Activity{
		Id:        a.id,
		StartedAt: timestamppb.New(a.startedAt),
		Duration:  durationpb.New(time.Since(a.startedAt)),
		Request:   anyToAnyPb(a.req),
		Props:     a.props,
		Events:    make([]*diveContracts.Activity_Event, 0, len(a.events)),
	}

	for _, e := range a.events {
		ret.Events = append(ret.Events, EventToProto(e))
	}

	return ret
}

func EventToProto(e Event) *diveContracts.Activity_Event {
	ret := &diveContracts.Activity_Event{
		Name:      e.name,
		Timestamp: timestamppb.New(e.timestamp),
		Props:     e.payload,
	}

	return ret
}

func ActivityFromProto(a *diveContracts.Activity) *Activity {
	ret := &Activity{
		id:        a.Id,
		startedAt: a.StartedAt.AsTime(),
		req:       a.Request,
		props:     a.Props,
		events:    make([]Event, 0, len(a.Events)),
	}

	for _, e := range a.Events {
		ret.events = append(ret.events, EventFromProto(e))
	}

	return ret
}

func EventFromProto(e *diveContracts.Activity_Event) Event {
	ret := Event{
		name:      e.Name,
		timestamp: e.Timestamp.AsTime(),
		payload:   e.Props,
	}

	return ret
}
