package person

import (
	"context"
	"distributor/workers/distributors"
	"fmt"

	"github.com/hibiken/asynq"
)

func CreateFidelityRegister(ctx context.Context, fidelityAmbassadorToken string) error {

	err := distributors.NewRedisTaskDistributor().DistributorCreditUserPoints(ctx, &distributors.PayloadCreditUserPoints{
		ReferralID:     22,
		ReferrerUserID: 22,
	}, asynq.MaxRetry(5))
	if err != nil {
		return fmt.Errorf("distributorCreditUserPoints -> %w", err)
	}

	return nil
}
