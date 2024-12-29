package person

import (
	"context"
	"grpc-caller/grpc/server/pb"
	"time"
)

type PersonGrpcHandlers struct {
	pb.UnimplementedPersonRPCServer
}

func (p *PersonGrpcHandlers) CreateFidelityRegister(ctx context.Context, in *pb.CreateFidelity) (*pb.ResponseDefault, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	response := &pb.ResponseDefault{}

	err := CreateFidelityRegister(ctx, in)
	if err != nil {
		response.Error = err.Error()
		return response, err
	}

	response.Message = "Process follows workflow normally..."

	return response, nil
}
