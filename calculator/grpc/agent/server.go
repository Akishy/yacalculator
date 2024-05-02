package agent

import (
	calcv1 "github.com/Akishy/yacalculator/proto/gen/calculator"
)

type serverAPI struct {
	calcv1.UnimplementedTaskServer
}

//func Register(gRPC *grpc.Server) {
//	calcv1.RegisterAgentServer(gRPC, &serverAPI{})
//}

//func (s *serverAPI) Create(ctx context.Context, req *calcv1.CreateRequest) (*calcv1.CreateResponse, error) {
//
//}
