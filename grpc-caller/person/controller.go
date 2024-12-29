package person

import (
	"context"
	"fmt"
	"grpc-caller/grpc/server/pb"
	"grpc-caller/workers/distributors"
	"log"

	"github.com/hibiken/asynq"
)

func CreateFidelityRegister(ctx context.Context, payload *pb.CreateFidelity) error {

	workerPayload := &distributors.PayloadCreditUserPoints{
		IndicatorID: payload.IndicatorId,
		WantRetry:   payload.WantRetry,
	}

	log.Println("Starting worker process...")

	err := distributors.NewRedisTaskDistributor().DistributorCreditUserPoints(ctx, workerPayload, asynq.MaxRetry(3))
	if err != nil {
		return fmt.Errorf("distributorCreditUserPoints -> %w", err)
	}

	return nil
}
