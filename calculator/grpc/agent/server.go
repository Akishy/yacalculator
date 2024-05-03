package agent

import (
	"context"
	calcv1 "github.com/Akishy/yacalculator/proto/gen/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type serverAPI struct {
	calcv1.UnimplementedTaskServer
	calcv1.UnimplementedAgentServer
}

func Register(gRPC *grpc.Server) {
	calcv1.RegisterAgentServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Create(ctx context.Context, req *calcv1.CreateRequest) (*calcv1.CreateResponse, error) {
	if req.GetOwnerId() <= 0 {
		err := status.Error(codes.InvalidArgument, "owner id is required")
		log.Println(err)
		return nil, err
	}
	return &calcv1.CreateResponse{}, nil
}
