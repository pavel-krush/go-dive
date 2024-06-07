package go_dive

import (
	"context"

	diveContracts "github.com/pavel-krush/dive/api/go/v1"
)

type GrpcServer struct {
	diveContracts.UnimplementedDiveServiceServer
	master *Master
}

func NewGrpcServer(master *Master) *GrpcServer {
	return &GrpcServer{master: master}
}

func (s *GrpcServer) Dive(ctx context.Context, req *diveContracts.DiveRequest) (*diveContracts.DiveResponse, error) {
	resp := &diveContracts.DiveResponse{}
	s.master.WalkActivities(func(a *Activity) {
		resp.Activities = append(resp.Activities, a.ToProto())
	})
	return resp, nil
}
