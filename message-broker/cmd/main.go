package main

import (
	"context"
	"log"
	"message-broker/config"
	"message-broker/workers"
	process "message-broker/workers/enums"
	"message-broker/workers/processors"
	"time"

	redis_worker "github.com/redis/go-redis/v9"

	logZero "github.com/rs/zerolog/log"

	"github.com/hibiken/asynq"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("config file was not loaded correctly: %v\n", err)
	}

	err = StartWorkers(cfg)
	if err != nil {
		log.Fatalf("worker didnt start: %v\n", err)
	}
}

func StartWorkers(cfg config.AppConfig) error {
	redisOpt := asynq.RedisClientOpt{Addr: cfg.RedisAdr, DB: cfg.RedisDB}

	logger := workers.NewLogger()

	redis_worker.SetLogger(logger)

	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Concurrency:     15,
			ShutdownTimeout: 30 * time.Minute,
			Queues:          map[string]int{process.ProcessCreditUserPoints.QueueName: 1},
			ErrorHandler: asynq.ErrorHandlerFunc(
				func(ctx context.Context, task *asynq.Task, err error) {
					logZero.Error().Err(err).Str("type", task.Type()).Bytes("payload", task.Payload()).Msg("create server failed")
				}),
			Logger: logger,
		},
	)

	processorCreditUserPoints := processors.NewRedisTaskProcessorCreditUserPoints(server)

	mux := asynq.NewServeMux()

	mux.HandleFunc(process.ProcessCreditUserPoints.TaskName, processorCreditUserPoints.ProcessTaskCreditUserPoints)

	err := server.Run(mux)
	if err != nil {
		logZero.Fatal().Err(err).Str("erro", "start workers").Msg(err.Error())
		return err
	}

	return nil
}
