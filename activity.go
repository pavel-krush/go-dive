package dive

import (
	"maps"
	"slices"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	mu        sync.RWMutex
	id        string
	startedAt time.Time
	req       any
	props     map[string]string
	events    Events
}

type Events []Event

type Event struct {
	name      string
	timestamp time.Time
	payload   map[string]string
}

const (
	PropBackgroundJobID = "bg_job_id"
	PropGrpcType        = "grpc_type"
	PropGrpcTypeUnary   = "unary"
	PropGrpcTypeStream  = "stream"
	PropGrpcService     = "grpc_service"
	PropGrpcMethod      = "grpc_method"
)

func newActivity(req any) *Activity {
	return &Activity{
		id:        uuid.New().String(),
		startedAt: time.Now(),
		req:       req,
		props:     make(map[string]string),
	}
}

func NewBackgroundActivity(req any, jobID string) *Activity {
	ret := newActivity(req)
	ret.props[PropBackgroundJobID] = jobID
	return ret
}

func NewUnaryGrpcActivity(service string, method string, req any) *Activity {
	ret := newActivity(req)
	ret.SetProp(PropGrpcType, PropGrpcTypeUnary)
	ret.SetProp(PropGrpcService, service)
	ret.SetProp(PropGrpcMethod, method)
	return ret
}

func NewStreamGrpcActivity(service string, method string, req any) *Activity {
	ret := newActivity(req)
	ret.SetProp(PropGrpcType, PropGrpcTypeStream)
	ret.SetProp(PropGrpcService, service)
	ret.SetProp(PropGrpcMethod, method)
	return ret
}

func (a *Activity) SetProps(props map[string]string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for k, v := range props {
		a.props[k] = v
	}
}

func (a *Activity) SetProp(key string, value string) {
	if a == nil {
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.props[key] = value
}

func (a *Activity) ID() string {
	return a.id
}

func (a *Activity) StartedAt() time.Time {
	return a.startedAt
}

func (a *Activity) Duration() time.Duration {
	return time.Since(a.startedAt)
}

func (a *Activity) SetRequest(req any) {
	if a == nil {
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.req = req
}

func (a *Activity) Request() any {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.req
}

func (a *Activity) Props() map[string]string {
	ret := make(map[string]string)
	maps.Copy(ret, a.props)
	return ret
}

func (a *Activity) AddEvent(name string, ts time.Time, payload map[string]string) {
	if a == nil {
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.events = append(a.events, Event{
		name:      name,
		timestamp: ts,
		payload:   payload,
	})
}

func (a *Activity) AddEventNow(name string, payload map[string]string) {
	a.AddEvent(name, time.Now(), payload)
}

func (a *Activity) Events() []Event {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return slices.Clone(a.events)
}

func (e Event) ID() string {
	return e.name
}

func (e Event) Timestamp() time.Time {
	return e.timestamp
}

func (e Event) Payload() map[string]string {
	ret := make(map[string]string)
	maps.Copy(ret, e.payload)
	return ret
}
