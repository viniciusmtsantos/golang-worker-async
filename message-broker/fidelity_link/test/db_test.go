package fidelitylink_test

import (
	"testing"

	fidelitylink "worker-demo/fidelity_link"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// Para correto funcionamento do teste, as constantes config.*_FILE devem ser alteradas para o caminho relativo à pasta 'test'

func TestSetCreditedPointsByReferralIDDBTx_Success(t *testing.T) {

	assert.NoError(t, fidelitylink.SetCreditedPointsByReferralIDDBTx(validReferralID, points, &sqlx.Tx{}), "O erro deveria ser nil")
}

func TestSetCreditedPointsByReferralIDDBTx_Error(t *testing.T) {
	testCases := []struct {
		name             string
		referralID       int64
		responseContains string
	}{
		{"Error_InvalidReferralID", invalidID, "nenhum registro de indicação"},
		{"Error_NotFoundReferralID", 2, "nenhum registro de indicação"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := fidelitylink.SetCreditedPointsByReferralIDDBTx(tc.referralID, 10, &sqlx.Tx{})

			if assert.Error(t, err, "O erro não deveria ser nil") {
				assert.Contains(t, err.Error(), tc.responseContains)
			}
		})
	}
}

func TestInsertErrorLogDBTx_Success(t *testing.T) {

	assert.NoError(t, fidelitylink.InsertErrorLogDBTx(fidelitylink.AmbassadorReferralLog{
		AmbassadorReferralID: validReferralID,
		RequestType:          "PUT",
		Request:              "",
		Response:             "error message",
	}, &sqlx.Tx{}), "O erro deveria ser nil")
}

func TestInsertErrorLogDBTx_Error(t *testing.T) {

	testCases := []struct {
		name       string
		referralID int64
	}{
		{"Error_InvalidReferralID", invalidID},
		{"Error_NotFoundReferralID", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Error(t, fidelitylink.InsertErrorLogDBTx(fidelitylink.AmbassadorReferralLog{
				AmbassadorReferralID: tc.referralID,
				RequestType:          "PUT",
				Request:              "",
				Response:             "error message",
			}, &sqlx.Tx{}), "O erro não deveria ser nil")
		})
	}
}
