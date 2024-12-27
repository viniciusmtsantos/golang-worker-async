package fidelitylink_test

import (
	"context"
	"encoding/json"
	"testing"

	fidelitylink "worker-demo/fidelity_link"
	"worker-demo/workers/distributor"
	"worker-demo/workers/processors"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
)

// TestProcessTaskCreditUserPoints_Success executada SetCreditedPointsByReferralID com 0 pontos para que a indicação registrada com os pontos possa ser executada com sucesso nos proximos testes
func TestProcessTaskCreditUserPoints_Success(t *testing.T) {

	payloadBytes, err := json.Marshal(distributor.PayloadCreditUserPoints{
		ReferralID:     validReferralID,
		ReferrerUserID: validUserID,
	})
	if err != nil {
		t.Fatalf("Error marshalling payload: %v", err)
	}

	opts := []asynq.Option{asynq.MaxRetry(5)}
	task := asynq.NewTask("task_type", payloadBytes, opts...)
	processor := processors.NewRedisTaskProcessorCreditUserPoints(&asynq.Server{})

	assert.NoError(t, processor.ProcessTaskCreditUserPoints(context.Background(), task))

	assert.NoError(t, fidelitylink.SetCreditedPointsByReferralID(validReferralID, 0), "O erro deveria ser nil")
}

// TestProcessTaskCreditUserPoints_Success executada SetCreditedPointsByReferralID com 0 pontos para que a indicação registrada com os pontos possa ser executada com sucesso nos proximos testes
func TestProcessTaskCreditUserPoints_Error(t *testing.T) {

	processor := processors.NewRedisTaskProcessorCreditUserPoints(&asynq.Server{})

	testCases := []struct {
		name           string
		referralID     int64
		ReferrerUserID int64
	}{
		{"Error_NotFoundReferralID", 2, validUserID},
		{"Error_NotFoundUserID", validReferralID, 50},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			payload := distributor.PayloadCreditUserPoints{
				ReferralID:     tc.referralID,
				ReferrerUserID: tc.ReferrerUserID,
			}

			payloadBytes, err := json.Marshal(payload)
			if err != nil {
				t.Fatalf("Error marshalling payload: %v", err)
			}

			task := asynq.NewTask("task_type", payloadBytes)

			assert.Error(t, processor.ProcessTaskCreditUserPoints(context.Background(), task), "O erro não deveria ser nil")
		})
	}
}
