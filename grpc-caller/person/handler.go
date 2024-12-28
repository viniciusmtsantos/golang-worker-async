package person

import (
	"context"
	"grpc-caller/grpc/server/pb"
	time_go "time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type PersonGrpcHandlers struct {
	pb.UnimplementedPersonRPCServer
}

func (p *PersonGrpcHandlers) CreateFidelityRegister(ctx context.Context, in *pb.CreateFidelity) (*pb.ResponseDefault, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time_go.Second)
	defer cancel()

	defer newrelic.FromContext(ctx).StartSegment("Handler > CreateFidelityRegister").End()

	response := &pb.ResponseDefault{}

	if err := CreateFidelityRegister(ctx, in.FidelityToken); err != nil {
		response.Error = err.Error()
		return response, err
	}

	return response, nil
}
