package processors

import (
	"context"
	"encoding/json"
	"fmt"
	fidelitylink "worker-demo/fidelity_link"
	"worker-demo/workers/distributor"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type TaskProcessorCreditUserPoints interface {
	ProcessTaskCreditUserPoints(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessorCreditUserPoints struct {
	server *asynq.Server
}

func NewRedisTaskProcessorCreditUserPoints(server *asynq.Server) TaskProcessorCreditUserPoints {
	return &RedisTaskProcessorCreditUserPoints{
		server: server,
	}
}

func (processor *RedisTaskProcessorCreditUserPoints) ProcessTaskCreditUserPoints(ctx context.Context, task *asynq.Task) error {
	var payload distributor.PayloadCreditUserPoints

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		errLog := fmt.Errorf("failed to unmarshal payload processor status credit points: %w", asynq.SkipRetry)
		return errLog
	}

	if payload.ReferralID == 0 || payload.ReferrerUserID == 0 {
		errLog := fmt.Errorf("invalid payload to process status credit points: %w", asynq.SkipRetry)
		return errLog
	}

	personIndicateID, err := fidelitylink.FindPersonIDByReferralID(payload.ReferralID)
	if err != nil {
		errLog := fmt.Errorf("error: %s - %v", err.Error(), payload)
		return errLog
	}

	externalID, err := fidelitylink.GetDocumentByUserID(personIndicateID)
	if err != nil {
		return err
	}

	points, err := fidelitylink.GetPointsByParameterName(fidelitylink.PointsParameterName)
	if err != nil {
		return err
	}

	err = fidelitylink.CreditPointsToReferrer(externalID, points)
	if err != nil {
		return err
	}

	err = fidelitylink.SetCreditedPointsByReferralID(payload.ReferralID, points)
	if err != nil {
		return err
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Int64("ReferralID", payload.ReferralID).Msg("processed task credit  user points")
	return nil
}
