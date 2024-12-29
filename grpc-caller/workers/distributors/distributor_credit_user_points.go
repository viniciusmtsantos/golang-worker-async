package distributors

import (
	"context"
	"encoding/json"
	"grpc-caller/config"
	"grpc-caller/workers/enums/process"

	"github.com/hibiken/asynq"
	log "github.com/sirupsen/logrus"
)

type RedisTaskDistributorCreditUserPoints struct {
	client *asynq.Client
}

type TaskDistributorCreditUserPoints interface {
	DistributorCreditUserPoints(ctx context.Context, payload *PayloadCreditUserPoints, opts ...asynq.Option) error
}

func NewRedisTaskDistributor() TaskDistributorCreditUserPoints {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("config file was not loaded correctly: %v\n", err)
	}

	redisOpt := asynq.RedisClientOpt{
		Addr: cfg.RedisAdr,
		DB:   cfg.RedisDB,
	}

	client := asynq.NewClient(redisOpt)

	return &RedisTaskDistributorCreditUserPoints{
		client: client,
	}
}

func (s *RedisTaskDistributorCreditUserPoints) DistributorCreditUserPoints(ctx context.Context, payload *PayloadCreditUserPoints, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	task := asynq.NewTask(process.ProcessCreditUserPoints.TaskName, jsonPayload, asynq.Queue(process.ProcessCreditUserPoints.QueueName))

	info, err := s.client.EnqueueContext(ctx, task)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"type":               task.Type(),
		"payload":            string(task.Payload()),
		"queue":              info.Queue,
		"max_retry_possible": info.MaxRetry,
	}).Info("enqueued task")

	return nil
}
