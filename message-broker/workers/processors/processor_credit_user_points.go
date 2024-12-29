package processors

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"message-broker/taskflow"
	"message-broker/workers/distributor"

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

	if payload.IndicatorID == 0 {
		errLog := fmt.Errorf("invalid payload to process status credit points: %w", asynq.SkipRetry)
		return errLog
	}

	points, err := taskflow.GetPointsByParameterName(taskflow.PointsParameterName)
	if err != nil || payload.WantRetry {
		return errors.New("retry exceptions testing")
	}

	err = taskflow.CreditPointsToReferrer(payload.IndicatorID, points)
	if err != nil {
		return err
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Int64("ReferralID", payload.IndicatorID).Msg("processed task credit  user points")

	return nil
}
